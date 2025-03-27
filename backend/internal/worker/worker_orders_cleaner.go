package worker

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/manager/transaction"
)

type WorkerOrdersCleaner struct {
	cfg               Config
	ordersRepo        OrdersRepository
	productCountsRepo ProductCountsRepository
	orderProductsRepo OrderProductsRepository
	txManager         *transaction.Manager

	wg *sync.WaitGroup
}

type Config interface {
	OrderMinutesTTL() int64
	OrdersCleanupIntervalMinutes() int64
}

type OrdersRepository interface {
	DeleteOrder(ctx context.Context, orderUid uuid.UUID) error
	GetOrderWithProducts(ctx context.Context, userUid uuid.UUID, qFilters entity.QueryFilters) ([]entity.OrderWithProducts, error)
}

type OrderProductsRepository interface {
	DeleteOrderProducts(ctx context.Context, orderUid uuid.UUID) error
}

type ProductCountsRepository interface {
	UpdateCount(ctx context.Context, productsCounts entity.ProductsCounts) error
}

func NewWorkerOrdersCleaner(
	cfg Config,
	ordersRepo OrdersRepository,
	productCountsRepo ProductCountsRepository,
	orderProductsRepo OrderProductsRepository,
	txManager *transaction.Manager,
) *WorkerOrdersCleaner {
	return &WorkerOrdersCleaner{
		cfg:               cfg,
		ordersRepo:        ordersRepo,
		productCountsRepo: productCountsRepo,
		orderProductsRepo: orderProductsRepo,
		txManager:         txManager,
		wg:                &sync.WaitGroup{},
	}
}

func (w *WorkerOrdersCleaner) Run(ctx context.Context) {
	(&sync.Once{}).Do(func() {
		fmt.Println("workerOrdersCleaner started")
		w.wg.Add(1)
		go func() {
			timer := time.NewTimer(time.Duration(w.cfg.OrdersCleanupIntervalMinutes()) * time.Minute)

			for {
				select {
				case <-ctx.Done():
					log.Printf("workerOrdersCleaner stopped: %v", ctx.Err())
					timer.Stop()
					return
				case <-timer.C:
					select {
					case <-ctx.Done():
						log.Printf("workerOrdersCleaner stopped: %v", ctx.Err())
						timer.Stop()
						return
					default:
					}

					if err := w.clean(ctx); err != nil {
						log.Printf("workerOrdersCleaner: failed to clean unused orders: %v", err)
					}

					timer.Reset(time.Duration(w.cfg.OrdersCleanupIntervalMinutes()) * time.Minute)
				}
			}
		}()
	})
}

func (w *WorkerOrdersCleaner) Close() {
	w.wg.Wait()
	log.Printf("workerOrdersCleaner successfully stopped")
}

func (w *WorkerOrdersCleaner) clean(ctx context.Context) error {
	updatedBefore := time.Now().Add(-time.Duration(w.cfg.OrderMinutesTTL()) * time.Minute)

	ordersInfo, err := w.ordersRepo.GetOrderWithProducts(ctx, uuid.Nil, entity.QueryFilters{
		OrderStatusNotIn: []string{string(entity.OrderStatusDone), string(entity.OrderStatusPaid)},
		UpdatedBefore:    updatedBefore,
	})
	if err != nil {
		log.Printf(
			"workerOrdersCleaner: failed to get orders with products with status not paid, done and updated before %s: %v",
			updatedBefore, err)
		return err
	}

	for _, orderInfo := range ordersInfo {

		time.Sleep(1 * time.Second) // пока так чтобы не нагнуть базу

		if err := w.txManager.NewPgTransaction().Execute(ctx, func(ctx context.Context) error {
			// взять заказ с продуктами
			infos, err := w.ordersRepo.GetOrderWithProducts(ctx, uuid.Nil, entity.QueryFilters{
				OrderStatusNotIn: []string{string(entity.OrderStatusDone), string(entity.OrderStatusPaid)},
				UpdatedBefore:    time.Now().Add(-time.Duration(w.cfg.OrderMinutesTTL()) * time.Minute),
				OrderUid:         orderInfo.Uid,
			})
			if err != nil {
				log.Printf(
					"workerOrdersCleaner: failed to get order %s with products with status not paid, done and updated before %s: %v",
					orderInfo.Uid, updatedBefore, err)
				return err
			}
			if len(infos) == 0 {
				// заказ перешел в следующую стадию, пропускаем
				log.Printf("workerOrdersCleaner: order %s turned to next state, skip order", orderInfo.Uid)
				return nil
			}

			// // удалить продукты  заказа
			// if err := w.orderProductsRepo.DeleteOrderProducts(ctx, infos[0].Order.Uid); err != nil {
			// 	log.Printf("workerOrdersCleaner: failed to delete order %s products: %v", orderInfo.Uid, err)
			// 	return err
			// }

			// // обновить счетчики
			// pCounts := make([]entity.ProductCount, 0, len(infos[0].Products))
			// for _, p := range infos[0].Products {
			// 	pCounts = append(pCounts, entity.ProductCount{
			// 		ProductUid: p.ProductUid,
			// 		Count:      -p.Count, // минус потому что в UpdateCount вычитание / нахуя названо Update тогда я хз
			// 	})
			// }
			// if err := w.productCountsRepo.UpdateCount(ctx, entity.ProductsCounts{
			// 	Products: pCounts,
			// }); err != nil {
			// 	log.Printf("workerOrdersCleaner: failed to return counts from order %s products: %v", orderInfo.Uid, err)
			// 	return err
			// }

			// // удалить заказ
			// if err := w.ordersRepo.DeleteOrder(ctx, infos[0].Order.Uid); err != nil {
			// 	log.Printf("workerOrdersCleaner: failed to delete order %s: %v", orderInfo.Uid, err)
			// 	return err
			// }
			// return nil

			return nil
		}); err != nil {
			log.Printf("failed to clean unused order %s: %v", orderInfo.Uid, err)
		}
	}
	return nil
}
