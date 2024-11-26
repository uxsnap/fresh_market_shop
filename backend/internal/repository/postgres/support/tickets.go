package repositorySupport

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *SupportRepository) CreateSupportTicket(ctx context.Context, ticket entity.SupportTicket) error {
	log.Printf("repositorySupport.CreateSupportTicket: title %s", ticket.Title)

	if err := r.Create(ctx, pgEntity.NewSupportTicketRow().FromEntity(ticket)); err != nil {
		log.Printf("failed to create support ticket: %v", err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *SupportRepository) UpdateSupportTicket(ctx context.Context, ticket entity.SupportTicket) error {
	log.Printf("repositorySupport.UpdateSupportTicket: uid %s", ticket.Uid)

	row := pgEntity.NewSupportTicketRow().FromEntity(ticket)
	if err := r.Update(ctx, row, row.ConditionUidEqual()); err != nil {
		log.Printf("failed to update support ticket %s: %v", ticket.Uid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *SupportRepository) DeleteSupportTicket(ctx context.Context, uid uuid.UUID) error {
	log.Printf("repositorySupport.DeleteSupportTicket: uid %s", uid)

	row := pgEntity.NewSupportTicketRow().FromEntity(entity.SupportTicket{Uid: uid})
	if err := r.Delete(ctx, row, row.ConditionUidEqual()); err != nil {
		log.Printf("failed to delete support ticket %s: %v", uid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *SupportRepository) GetSupportTicketByUid(ctx context.Context, uid uuid.UUID) (entity.SupportTicket, bool, error) {
	log.Printf("repositorySupport.GetSupportTicketByUid: uid %s", uid)

	row := pgEntity.NewSupportTicketRow().FromEntity(entity.SupportTicket{Uid: uid})
	if err := r.GetOne(ctx, row, row.ConditionUidEqual()); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.SupportTicket{}, false, nil
		}
		log.Printf("failed to get support ticket %s: %v", uid, err)
		return entity.SupportTicket{}, false, errors.WithStack(err)
	}
	return row.ToEntity(), true, nil
}

func (r *SupportRepository) GetSupportTickets(ctx context.Context, qFilters entity.QueryFilters) ([]entity.SupportTicket, error) {
	log.Printf(
		"repositorySupport.GetSupportTickets: user_uid %s , topic_uid %s, solver_uid %s, from_email %s, limit %d, offset %d",
		qFilters.UserUid, qFilters.TopicUid, qFilters.SolverUid, qFilters.FromEmail, qFilters.Limit, qFilters.Offset,
	)
	row := pgEntity.NewSupportTicketRow().FromEntity(entity.SupportTicket{
		UserUid:   qFilters.UserUid,
		TopicUid:  qFilters.TopicUid,
		SolverUid: qFilters.SolverUid,
		FromEmail: qFilters.FromEmail,
		Status:    entity.SupportTicketStatus(qFilters.Status),
	})

	conds := make([]sq.Sqlizer, 0, 5)
	if !uuid.Equal(qFilters.UserUid, uuid.UUID{}) {
		conds = append(conds, row.ConditionUserUidEqual())
	}
	if !uuid.Equal(qFilters.TopicUid, uuid.UUID{}) {
		conds = append(conds, row.ConditionTopicUidEqual())
	}
	if !uuid.Equal(qFilters.SolverUid, uuid.UUID{}) {
		conds = append(conds, row.ConditionSolverUidEqual())
	}
	if len(qFilters.FromEmail) != 0 {
		conds = append(conds, row.ConditionFromEmailEqual())
	}
	if len(qFilters.Status) != 0 {
		conds = append(conds, row.ConditionStatusEqual())
	}
	cond := append(sq.And{}, conds...)

	rows := pgEntity.NewSupportTicketRows()
	var err error

	if qFilters.Limit != 0 {
		err = r.GetWithLimit(ctx, row, rows, cond, qFilters.Limit, qFilters.Offset)
	} else {
		err = r.GetSome(ctx, row, rows, cond)
	}
	if err != nil {
		log.Printf("failed to get support tickets: %v", err)
		return nil, errors.WithStack(err)
	}

	return rows.ToEntity(), nil
}
