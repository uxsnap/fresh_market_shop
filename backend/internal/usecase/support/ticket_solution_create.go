package useCaseSupport

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseSupport) CreateTicketSolution(ctx context.Context, solution entity.SupportTicketSolution) error {
	log.Printf("usecaseSupport.CreateTicketSolution: ticketUid %s", solution.TicketUid)

	if uuid.Equal(solution.TicketUid, uuid.UUID{}) {
		log.Printf("failed to create ticket solution: empty ticket uid")
		return errors.New("пустой uid обращения")
	}
	if len(solution.Description) == 0 {
		log.Printf("failed to create ticket solution: empty solution description")
		return errors.New("пустое описание решения")
	}

	if len(solution.EmailText) == 0 {
		solution.EmailText = solution.Description
	}
	solution.CreatedAt = time.Now().UTC()

	if err := uc.txManager.NewPgTransaction().Execute(ctx, func(ctx context.Context) error {
		ticket, isFound, err := uc.repository.GetSupportTicketByUid(ctx, solution.TicketUid)
		if err != nil {
			log.Printf("failed to create ticket (%s) solution: %v", solution.TicketUid, err)
			return errors.WithStack(err)
		}
		if !isFound {
			log.Printf("failed to create ticket (%s) solution: ticket not found", solution.TicketUid)
			return errors.New("обращение не найдено")
		}

		if solution.IsSuccess {
			ticket.Status = entity.SupportTicketStatusSolved
		} else {
			ticket.Status = entity.SupportTicketStatusCantSolve
		}

		if err := uc.repository.UpdateSupportTicket(ctx, ticket); err != nil {
			log.Printf("failed to create ticket (%s) solution: %v", ticket.Uid, err)
			return errors.WithStack(err)
		}
		if err := uc.repository.CreateSupportTicketSolution(ctx, solution); err != nil {
			log.Printf("failed to create ticket (%s) solution: %v", ticket.Uid, err)
			return errors.WithStack(err)
		}
		// TODO: отправлять на почту уведомление (outbox). Брать ticket.FromEmail в качестве получателя
		return nil
	}); err != nil {
		return err
	}

	return nil
}
