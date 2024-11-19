package app

import (
	"context"
	"log"

	DBclient "github.com/uxsnap/fresh_market_shop/backend/internal/client/database"
	"github.com/uxsnap/fresh_market_shop/backend/internal/client/database/pg"
	clientAuthService "github.com/uxsnap/fresh_market_shop/backend/internal/client/services/auth"
	"github.com/uxsnap/fresh_market_shop/backend/internal/config"
	deliveryHttp "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http"
	"github.com/uxsnap/fresh_market_shop/backend/internal/manager/transaction"
	repositoryAddresses "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/addresses"
	repositoryCategories "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/categories"
	repositoryDelivery "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/delivery"
	repositoryOrderProducts "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/order_products"
	repositoryOrders "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/orders"
	repositoryPayments "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/payments"
	repositoryProducts "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/products"
	repositoryProductsCount "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/products_count"
	repositoryRecipes "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/recipes"
	repositoryUsers "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/users"
	ucAddresses "github.com/uxsnap/fresh_market_shop/backend/internal/usecase/addresses"
	ucDelivery "github.com/uxsnap/fresh_market_shop/backend/internal/usecase/delivery"
	ucOrders "github.com/uxsnap/fresh_market_shop/backend/internal/usecase/orders"
	ucPayments "github.com/uxsnap/fresh_market_shop/backend/internal/usecase/payments"
	ucProducts "github.com/uxsnap/fresh_market_shop/backend/internal/usecase/products"
	ucRecipes "github.com/uxsnap/fresh_market_shop/backend/internal/usecase/recipes"
	ucUsers "github.com/uxsnap/fresh_market_shop/backend/internal/usecase/users"
)

type serviceProvider struct {
	configPG          *config.ConfigPG
	configHTTP        *config.ConfigHTTP
	configExternalApi *config.ConfigExternalApi

	pgClient   DBclient.ClientDB
	authClient *clientAuthService.AuthClient

	productsRepository      *repositoryProducts.ProductsRepository
	categoriesRepository    *repositoryCategories.CategoriesRepository
	usersRepository         *repositoryUsers.UsersRepository
	recipesRepository       *repositoryRecipes.RecipesRepository
	ordersRepository        *repositoryOrders.OrdersRepository
	productsCountRepository *repositoryProductsCount.ProductsCountRepository
	orderProductsRepository *repositoryOrderProducts.OrderProductsRepository
	deliveryRepository      *repositoryDelivery.DeliveryRepository
	addressesRepository     *repositoryAddresses.AddressesRepository
	paymentsRepository      *repositoryPayments.PaymentsRepository

	ucProducts  *ucProducts.UseCaseProducts
	ucUsers     *ucUsers.UseCaseUsers
	ucRecipes   *ucRecipes.UseCaseRecipes
	ucOrders    *ucOrders.UseCaseOrders
	ucDelivery  *ucDelivery.UseCaseDelivery
	ucAddresses *ucAddresses.UseCaseAddresses
	ucPayments  *ucPayments.UseCasePayments

	txManager *transaction.Manager

	handlerHTTP *deliveryHttp.Handler
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) ConfigPG() *config.ConfigPG {
	if sp.configPG == nil {
		sp.configPG = config.NewConfigPG()
	}
	return sp.configPG
}

func (sp *serviceProvider) ConfigHTTP() *config.ConfigHTTP {
	if sp.configHTTP == nil {
		sp.configHTTP = config.NewConfigHTTP()
	}
	return sp.configHTTP
}

func (sp *serviceProvider) ConfigExternalApi() *config.ConfigExternalApi {
	if sp.configExternalApi == nil {
		sp.configExternalApi = config.NewConfigExternalApi()
	}
	return sp.configExternalApi
}

func (sp *serviceProvider) PgClient(ctx context.Context) DBclient.ClientDB {
	if sp.pgClient == nil {
		client, err := pg.NewClient(ctx, sp.ConfigPG().DSN())
		if err != nil {
			log.Fatalf("failed to connect to postgres: %v", err)
		}
		sp.pgClient = client
	}
	return sp.pgClient
}

func (sp *serviceProvider) TxManager(ctx context.Context) *transaction.Manager {
	if sp.txManager == nil {
		sp.txManager = transaction.NewTxManager(sp.PgClient(ctx))
	}
	return sp.txManager
}

func (sp *serviceProvider) AuthClient(ctx context.Context) *clientAuthService.AuthClient {
	if sp.authClient == nil {
		client, err := clientAuthService.New(ctx, sp.ConfigExternalApi())
		if err != nil {
			log.Fatalf("failed to create auth service client: %v", err)
		}
		sp.authClient = client
	}
	return sp.authClient
}

func (sp *serviceProvider) ProductsRepository(ctx context.Context) *repositoryProducts.ProductsRepository {
	if sp.productsRepository == nil {
		sp.productsRepository = repositoryProducts.New(sp.PgClient(ctx))
	}
	return sp.productsRepository
}

func (sp *serviceProvider) CategoriesRepository(ctx context.Context) *repositoryCategories.CategoriesRepository {
	if sp.categoriesRepository == nil {
		sp.categoriesRepository = repositoryCategories.New(sp.PgClient(ctx))
	}
	return sp.categoriesRepository
}

func (sp *serviceProvider) UsersRepository(ctx context.Context) *repositoryUsers.UsersRepository {
	if sp.usersRepository == nil {
		sp.usersRepository = repositoryUsers.New(sp.PgClient(ctx))
	}
	return sp.usersRepository
}

func (sp *serviceProvider) RecipesRepository(ctx context.Context) *repositoryRecipes.RecipesRepository {
	if sp.recipesRepository == nil {
		sp.recipesRepository = repositoryRecipes.New(sp.PgClient(ctx))
	}
	return sp.recipesRepository
}

func (sp *serviceProvider) OrdersRepository(ctx context.Context) *repositoryOrders.OrdersRepository {
	if sp.ordersRepository == nil {
		sp.ordersRepository = repositoryOrders.New(sp.PgClient(ctx))
	}
	return sp.ordersRepository
}

func (sp *serviceProvider) ProductsCountRepository(ctx context.Context) *repositoryProductsCount.ProductsCountRepository {
	if sp.productsCountRepository == nil {
		sp.productsCountRepository = repositoryProductsCount.New(sp.PgClient(ctx))
	}
	return sp.productsCountRepository
}

func (sp *serviceProvider) OrderProductsRepository(ctx context.Context) *repositoryOrderProducts.OrderProductsRepository {
	if sp.orderProductsRepository == nil {
		sp.orderProductsRepository = repositoryOrderProducts.New(sp.PgClient(ctx))
	}
	return sp.orderProductsRepository
}

func (sp *serviceProvider) DeliveryRepository(ctx context.Context) *repositoryDelivery.DeliveryRepository {
	if sp.deliveryRepository == nil {
		sp.deliveryRepository = repositoryDelivery.New(sp.PgClient(ctx))
	}
	return sp.deliveryRepository
}

func (sp *serviceProvider) AddressesRepository(ctx context.Context) *repositoryAddresses.AddressesRepository {
	if sp.addressesRepository == nil {
		sp.addressesRepository = repositoryAddresses.New(sp.PgClient(ctx))
	}
	return sp.addressesRepository
}

func (sp *serviceProvider) PaymentsRepository(ctx context.Context) *repositoryPayments.PaymentsRepository {
	if sp.paymentsRepository == nil {
		sp.paymentsRepository = repositoryPayments.New(sp.PgClient(ctx))
	}
	return sp.paymentsRepository
}

func (sp *serviceProvider) ProductsService(ctx context.Context) *ucProducts.UseCaseProducts {
	if sp.ucProducts == nil {
		sp.ucProducts = ucProducts.New(
			sp.ProductsRepository(ctx),
			sp.CategoriesRepository(ctx),
			sp.TxManager(ctx),
		)
	}
	return sp.ucProducts
}

func (sp *serviceProvider) UsersService(ctx context.Context) *ucUsers.UseCaseUsers {
	if sp.ucUsers == nil {
		sp.ucUsers = ucUsers.New(sp.UsersRepository(ctx))
	}
	return sp.ucUsers
}

func (sp *serviceProvider) RecipesService(ctx context.Context) *ucRecipes.UseCaseRecipes {
	if sp.ucRecipes == nil {
		sp.ucRecipes = ucRecipes.New(sp.RecipesRepository(ctx), sp.ProductsRepository(ctx))
	}
	return sp.ucRecipes
}

func (sp *serviceProvider) OrdersService(ctx context.Context) *ucOrders.UseCaseOrders {
	if sp.ucOrders == nil {
		sp.ucOrders = ucOrders.New(
			sp.OrdersRepository(ctx),
			sp.OrderProductsRepository(ctx),
			sp.ProductsCountRepository(ctx),
			sp.ProductsRepository(ctx),
			sp.PaymentsService(ctx),
			sp.DeliveryService(ctx),
			sp.TxManager(ctx),
		)
	}
	return sp.ucOrders
}

func (sp *serviceProvider) DeliveryService(ctx context.Context) *ucDelivery.UseCaseDelivery {
	if sp.ucDelivery == nil {
		sp.ucDelivery = ucDelivery.New(
			sp.DeliveryRepository(ctx),
			sp.UsersService(ctx),
			sp.TxManager(ctx),
		)
	}
	return sp.ucDelivery
}

func (sp *serviceProvider) AddressesService(ctx context.Context) *ucAddresses.UseCaseAddresses {
	if sp.ucAddresses == nil {
		sp.ucAddresses = ucAddresses.New(
			sp.AddressesRepository(ctx),
			sp.TxManager(ctx),
		)
	}
	return sp.ucAddresses
}

func (sp *serviceProvider) PaymentsService(ctx context.Context) *ucPayments.UseCasePayments {
	if sp.ucPayments == nil {
		sp.ucPayments = ucPayments.New(
			sp.PaymentsRepository(ctx),
			sp.UsersService(ctx),
			sp.TxManager(ctx),
		)
	}
	return sp.ucPayments
}

func (sp *serviceProvider) HandlerHTTP(ctx context.Context) *deliveryHttp.Handler {
	if sp.handlerHTTP == nil {
		sp.handlerHTTP = deliveryHttp.New(
			nil,
			sp.AuthClient(ctx),
			sp.ProductsService(ctx),
			sp.UsersService(ctx),
			sp.RecipesService(ctx),
			sp.OrdersService(ctx),
			sp.DeliveryService(ctx),
			sp.AddressesService(ctx),
			sp.PaymentsService(ctx),
		)
	}
	return sp.handlerHTTP
}
