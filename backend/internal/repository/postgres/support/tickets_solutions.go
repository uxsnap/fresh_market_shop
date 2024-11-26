package repositorySupport

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *SupportRepository) CreateSupportTicketSolution(ctx context.Context, solution entity.SupportTicketSolution) error {
	log.Printf("repositorySupport.CreateSupportTicketSolution: ticket uid %s", solution.TicketUid)

	if err := r.Create(ctx, pgEntity.NewSupportTicketSolutionRow().FromEntity(solution)); err != nil {
		log.Printf("failed to create support ticket solution for ticket %s: %v", solution.TicketUid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *SupportRepository) UpdateSupportTicketSolution(ctx context.Context, solution entity.SupportTicketSolution) error {
	log.Printf("repositorySupport.UpdateSupportTicketSolution: ticket uid %s", solution.TicketUid)

	row := pgEntity.NewSupportTicketSolutionRow().FromEntity(solution)
	if err := r.Update(ctx, row, row.ConditionTicketUidEqual()); err != nil {
		log.Printf("failed to create support ticket solution for ticket %s: %v", solution.TicketUid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *SupportRepository) GetSupportTicketSolution(ctx context.Context, ticketUid uuid.UUID) (entity.SupportTicketSolution, bool, error) {
	log.Printf("repositorySupport.GetSupportTicketSolution: ticket uid %s", ticketUid)

	row := pgEntity.NewSupportTicketSolutionRow().FromEntity(entity.SupportTicketSolution{TicketUid: ticketUid})
	if err := r.GetOne(ctx, row, row.ConditionTicketUidEqual()); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.SupportTicketSolution{}, false, nil
		}
		log.Printf("failed to get support ticket solution with ticket uid %s: %v", ticketUid, err)
		return entity.SupportTicketSolution{}, false, errors.WithStack(err)
	}
	return row.ToEntity(), true, nil
}

func (r *SupportRepository) GetSupportTicketSolutionsByTopic(ctx context.Context, qFilters entity.QueryFilters) ([]entity.SupportTicketSolution, error) {
	log.Printf("repositorySupport.GetSupportTicketSolutionsByTopic: topic uid %s", qFilters.TopicUid)

	if uuid.Equal(qFilters.TopicUid, uuid.UUID{}) {
		return nil, errors.New("empty topic_uid")
	}

	row := pgEntity.NewSupportTicketSolutionRow().FromEntity(entity.SupportTicketSolution{TicketUid: qFilters.TopicUid})

	sql := sq.Select(
		withPrefix("s", row.Columns())...,
	).From(
		row.Table() + " s",
	).Join(
		pgEntity.NewSupportTicketRow().Table() + " t on s.ticket_uid=t.uid",
	).Where(
		sq.Eq{
			"t.topic_uid": pgtype.UUID{Bytes: qFilters.TopicUid, Status: pgtype.Present},
		},
	)
	if qFilters.Limit != 0 {
		sql = sql.Limit(qFilters.Limit).Offset(qFilters.Offset)
	}

	stmt, args, err := sql.ToSql()
	if err != nil {
		log.Printf("failed to build sql query for get support tickets solutions: %v", err)
		return nil, errors.WithStack(err)
	}

	rows, err := r.DB().Query(ctx, stmt, args...)
	if err != nil {
		log.Printf("failed to get support tickets solutions: %v", err)
		return nil, errors.WithStack(err)
	}

	solutionsRows := pgEntity.NewSupportTicketSolutionRows()
	if err := solutionsRows.ScanAll(rows); err != nil {
		log.Printf("failed to scan support tickets solutions: %v", err)
		return nil, errors.WithStack(err)
	}

	return solutionsRows.ToEntity(), nil
}

func (r *SupportRepository) DeleteSupportTicketSolution(ctx context.Context, ticketUid uuid.UUID) error {
	log.Printf("repositorySupport.DeleteSupportTicketSolution: ticket uid %s", ticketUid)

	row := pgEntity.NewSupportTicketSolutionRow().FromEntity(entity.SupportTicketSolution{TicketUid: ticketUid})
	if err := r.Delete(ctx, row, row.ConditionTicketUidEqual()); err != nil {
		log.Printf("failed to delete support ticket solution %s: %v", ticketUid, err)
		return errors.WithStack(err)
	}
	return nil
}

func withPrefix(prefix string, fields []string) []string {
	res := make([]string, 0, len(fields))
	for _, f := range fields {
		res = append(res, prefix+"."+f)
	}
	return res
}
