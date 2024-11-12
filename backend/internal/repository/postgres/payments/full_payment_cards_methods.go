package repositoryPayments

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *PaymentsRepository) CreateUserFullPaymentCard(ctx context.Context, card entity.UserFullPaymentCard) error {
	log.Printf("paymentsRepository.CreateUserFullPaymentCard: user uid %s", card.UserUid)

	cardRow := pgEntity.NewUserFullPaymentCardRow().FromEntity(card)
	if err := r.Create(ctx, cardRow); err != nil {
		log.Printf("failed to create payment card for user %s: %v", card.UserUid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *PaymentsRepository) GetUserFullPaymentCardByUid(ctx context.Context, uid uuid.UUID) (entity.UserFullPaymentCard, bool, error) {
	log.Printf("paymentsRepository.GetUserFullPaymentCardByUid: %s", uid)

	cardRow := pgEntity.NewUserFullPaymentCardRow().FromEntity(entity.UserFullPaymentCard{Uid: uid})
	if err := r.GetOne(ctx, cardRow, cardRow.ConditionUidEqual()); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.UserFullPaymentCard{}, false, nil
		}
		log.Printf("failed to get user full payment card by uid %s: %v", uid, err)
		return entity.UserFullPaymentCard{}, false, errors.WithStack(err)
	}
	return cardRow.ToEntity(), true, nil
}

func (r *PaymentsRepository) GetUserFullPaymentCards(ctx context.Context, userUid uuid.UUID) ([]entity.UserFullPaymentCard, error) {
	log.Printf("paymentsRepository.GetUserFullPaymentCards: user uid %s", userUid)

	cardRow := pgEntity.NewUserFullPaymentCardRow().FromEntity(entity.UserFullPaymentCard{UserUid: userUid})
	cardRows := pgEntity.NewUserFullPaymentCardRows()

	if err := r.GetSome(ctx, cardRow, cardRows, cardRow.ConditionUserUidEqual()); err != nil {
		log.Printf("failed to get payment cards by user uid %s: %v", userUid, err)
		return nil, errors.WithStack(err)
	}
	return cardRows.ToEntity(), nil
}

func (r *PaymentsRepository) DeleteUserFullPaymentCardByUid(ctx context.Context, cardUid uuid.UUID) error {
	log.Printf("paymentsRepository.DeleteUserFullPaymentCardByUid: card uid %s", cardUid)

	cardRow := pgEntity.NewUserFullPaymentCardRow().FromEntity(entity.UserFullPaymentCard{Uid: cardUid})

	if err := r.Delete(ctx, cardRow, cardRow.ConditionUidEqual()); err != nil {
		log.Printf("failed to delete full payment card by uid %s: %v", cardUid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *PaymentsRepository) DeleteUserFullPaymentCards(ctx context.Context, userUid uuid.UUID) error {
	log.Printf("paymentsRepository.DeleteUserFullPaymentCards: user uid %s", userUid)

	cardRow := pgEntity.NewUserFullPaymentCardRow().FromEntity(entity.UserFullPaymentCard{UserUid: userUid})

	if err := r.Delete(ctx, cardRow, cardRow.ConditionUserUidEqual()); err != nil {
		log.Printf("failed to delete full payment cards by user uid %s: %v", userUid, err)
		return errors.WithStack(err)
	}
	return nil
}
