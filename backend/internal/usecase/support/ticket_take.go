package useCaseSupport

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseSupport) TakeTicket(ctx context.Context, ticketUid, solverUid uuid.UUID) error {
	log.Printf("usecaseSupport.TakeTicket: uid %s, solver_uid %s", ticketUid, solverUid)

	if uuid.Equal(ticketUid, uuid.UUID{}) {
		return errors.New("пустой uid обращения")
	}
	if uuid.Equal(solverUid, uuid.UUID{}) {
		return errors.New("пустой uid менеджера поддержки")
	}

	if err := uc.txManager.NewPgTransaction().Execute(ctx, func(ctx context.Context) error {
		savedTicket, isFound, err := uc.repository.GetSupportTicketByUid(ctx, ticketUid)
		if err != nil {
			return err
		}
		if !isFound {
			return errors.New("обращение не найдено")
		}
		switch savedTicket.Status {
		case entity.SupportTicketStatusSolved, entity.SupportTicketStatusCantSolve:
			log.Printf("failed to take support ticket (%s) message: cant take ticket in status '%s'", savedTicket.Uid, savedTicket.Status)
			return errors.Errorf("нельзя взять в работу тикет в статусе '%s'", savedTicket.Status)
		default:
		}

		savedTicket.SolverUid = solverUid
		savedTicket.Status = entity.SupportTicketStatusInProcess
		savedTicket.UpdatedAt = time.Now().UTC()

		if err := uc.repository.UpdateSupportTicket(ctx, savedTicket); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Printf("failed to take ticket %s by support manager %s: %v", ticketUid, solverUid, err)
		return errors.WithStack(err)
	}
	return nil
}
