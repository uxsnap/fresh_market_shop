package repositorySupport

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *SupportRepository) CreateSupportTicketCommentMessage(ctx context.Context, message entity.SupportTicketCommentMessage) error {
	log.Printf("repositorySupport.CreateSupportTicketCommentMessage: ticket uid %s sender uid %s", message.TicketUid, message.SenderUid)

	if err := r.Create(ctx, pgEntity.NewSupportTicketCommentMessageRow().FromEntity(message)); err != nil {
		log.Printf("failed to create support ticket comment message for ticket %s: %v", message.TicketUid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *SupportRepository) UpdateSupportTicketCommentMessage(ctx context.Context, message entity.SupportTicketCommentMessage) error {
	log.Printf("repositorySupport.UpdateSupportTicketCommentMessage: ticket uid %s message uid %s", message.TicketUid, message.Uid)

	row := pgEntity.NewSupportTicketCommentMessageRow().FromEntity(message)
	if err := r.Update(ctx, row, row.ConditionUidEqual()); err != nil {
		log.Printf("failed to update support ticket comment message %s: %v", message.Uid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *SupportRepository) GetSupportTicketCommentMessage(ctx context.Context, uid uuid.UUID) (entity.SupportTicketCommentMessage, bool, error) {
	log.Printf("repositorySupport.GetSupportTicketCommentMessage: message uid %s", uid)

	row := pgEntity.NewSupportTicketCommentMessageRow().FromEntity(entity.SupportTicketCommentMessage{Uid: uid})
	if err := r.GetOne(ctx, row, row.ConditionUidEqual()); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.SupportTicketCommentMessage{}, false, nil
		}
		log.Printf("failed to get support ticket comment message %s: %v", uid, err)
		return entity.SupportTicketCommentMessage{}, false, errors.WithStack(err)
	}
	return row.ToEntity(), true, nil
}

func (r *SupportRepository) GetSupportTicketCommentMessages(ctx context.Context, ticketUid uuid.UUID, qFilters entity.QueryFilters) ([]entity.SupportTicketCommentMessage, error) {
	log.Printf("repositorySupport.GetSupportTicketCommentMessages: ticketUid %s", ticketUid)

	row := pgEntity.NewSupportTicketCommentMessageRow().FromEntity(entity.SupportTicketCommentMessage{TicketUid: ticketUid})
	rows := pgEntity.NewSupportTicketCommentMessageRows()
	if err := r.GetSome(ctx, row, rows, row.ConditionTicketUidEqual()); err != nil {
		log.Printf("failed to get support ticket (%s) comment messages: %v", ticketUid, err)
		return nil, errors.WithStack(err)
	}
	return rows.ToEntity(), nil
}
