package useCaseSupport

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseSupport) TakeTicket(ctx context.Context, ticket entity.SupportTicket) error {
	log.Printf("usecaseSupport.TakeTicket: uid %s, solver_uid %s", ticket.Uid, ticket.SolverUid)

	if uuid.Equal(ticket.Uid, uuid.UUID{}) {
		return errors.New("пустой uid обращения")
	}
	if uuid.Equal(ticket.SolverUid, uuid.UUID{}) {
		return errors.New("пустой uid менеджера поддержки")
	}

	if err := uc.txManager.NewPgTransaction().Execute(ctx, func(ctx context.Context) error {
		savedTicket, isFound, err := uc.repository.GetSupportTicketByUid(ctx, ticket.Uid)
		if err != nil {
			return err
		}
		if !isFound {
			return errors.New("обращение не найдено")
		}
		savedTicket.SolverUid = ticket.SolverUid
		savedTicket.Status = entity.SupportTicketStatusInProcess
		savedTicket.UpdatedAt = time.Now().UTC()

		if err := uc.repository.UpdateSupportTicket(ctx, savedTicket); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Printf("failed to take ticket %s by support manager %s: %v", ticket.Uid, ticket.SolverUid, err)
		return errors.WithStack(err)
	}
	return nil
}
