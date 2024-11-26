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

func (r *SupportRepository) CreateSupportTicketsTopic(ctx context.Context, topic entity.SupportTicketsTopic) error {
	log.Printf("repositorySupport.CreateSupportTicketsTopic: name %s", topic.Name)

	if err := r.Create(ctx, pgEntity.NewSupportTicketsTopicRow().FromEntity(topic)); err != nil {
		log.Printf("failed to create support tickets topic: %v", err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *SupportRepository) UpdateSupportTicketsTopic(ctx context.Context, topic entity.SupportTicketsTopic) error {
	log.Printf("repositorySupport.UpdateSupportTicketsTopic: name %s", topic.Name)
	row := pgEntity.NewSupportTicketsTopicRow().FromEntity(topic)

	if err := r.Update(ctx, row, row.ConditionUidEqual()); err != nil {
		log.Printf("failed to update support tickets topic %s: %v", topic.Uid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *SupportRepository) DeleteSupportTicketsTopic(ctx context.Context, uid uuid.UUID) error {
	log.Printf("repositorySupport.DeleteSupportTicketsTopic: uid %s", uid)
	row := pgEntity.NewSupportTicketsTopicRow().FromEntity(entity.SupportTicketsTopic{Uid: uid})

	if err := r.Delete(ctx, row, row.ConditionUidEqual()); err != nil {
		log.Printf("failed to delete support tickets topic %s: %v", uid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *SupportRepository) GetSupportTicketsTopicByUid(ctx context.Context, uid uuid.UUID) (entity.SupportTicketsTopic, bool, error) {
	log.Printf("repositorySupport.GetSupportTicketsTopicByUid: %s", uid)
	row := pgEntity.NewSupportTicketsTopicRow().FromEntity(entity.SupportTicketsTopic{Uid: uid})

	if err := r.GetOne(ctx, row, row.ConditionUidEqual()); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.SupportTicketsTopic{}, false, nil
		}
		log.Printf("failed to get support tickets topic %s: %v", uid, err)
		return entity.SupportTicketsTopic{}, false, errors.WithStack(err)
	}
	return row.ToEntity(), true, nil
}

func (r *SupportRepository) GetAllSupportTicketsTopics(ctx context.Context) ([]entity.SupportTicketsTopic, error) {
	log.Printf("repositorySupport.GetAllSupportTicketsTopics")

	row := pgEntity.NewSupportTicketsTopicRow()
	rows := pgEntity.NewSupportTicketsTopicRows()

	if err := r.GetSome(ctx, row, rows, nil); err != nil {
		log.Printf("failed to get all support tickets topics: %v", err)
		return nil, errors.WithStack(err)
	}
	return rows.ToEntity(), nil
}
