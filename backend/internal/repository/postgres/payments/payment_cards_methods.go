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

func (r *PaymentsRepository) CreateUserPaymentCard(ctx context.Context, card entity.UserPaymentCard) error {
	log.Printf("paymentsRepository.CreateUserPaymentCard: user uid %s", card.UserUid)

	cardRow := pgEntity.NewUserPaymentCardRow().FromEntity(card)
	if err := r.Create(ctx, cardRow); err != nil {
		log.Printf("failed to create payment card for user %s: %v", card.UserUid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *PaymentsRepository) GetUserPaymentCardByUid(ctx context.Context, uid uuid.UUID) (entity.UserPaymentCard, bool, error) {
	log.Printf("paymentsRepository.GetUserPaymentCardByUid: %s", uid)

	cardRow := pgEntity.NewUserPaymentCardRow().FromEntity(entity.UserPaymentCard{Uid: uid})
	if err := r.GetOne(ctx, cardRow, cardRow.ConditionUidEqual()); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.UserPaymentCard{}, false, nil
		}
		log.Printf("failed to get user payment card by uid %s: %v", uid, err)
		return entity.UserPaymentCard{}, false, errors.WithStack(err)
	}
	return cardRow.ToEntity(), true, nil
}

func (r *PaymentsRepository) GetUserPaymentCards(ctx context.Context, userUid uuid.UUID) ([]entity.UserPaymentCard, error) {
	log.Printf("paymentsRepository.GetUserPaymentCards: user uid %s", userUid)

	cardRow := pgEntity.NewUserPaymentCardRow().FromEntity(entity.UserPaymentCard{UserUid: userUid})
	cardRows := pgEntity.NewUserPaymentCardRows()

	if err := r.GetSome(ctx, cardRow, cardRows, cardRow.ConditionUserUidEqual()); err != nil {
		log.Printf("failed to get payment cards by user uid %s: %v", userUid, err)
		return nil, errors.WithStack(err)
	}
	return cardRows.ToEntity(), nil
}

func (r *PaymentsRepository) DeleteUserPaymentCardByUid(ctx context.Context, cardUid uuid.UUID) error {
	log.Printf("paymentsRepository.DeleteUserPaymentCardByUid: card uid %s", cardUid)

	cardRow := pgEntity.NewUserPaymentCardRow().FromEntity(entity.UserPaymentCard{Uid: cardUid})

	if err := r.Delete(ctx, cardRow, cardRow.ConditionUidEqual()); err != nil {
		log.Printf("failed to delete payment card by uid %s: %v", cardUid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *PaymentsRepository) DeleteUserPaymentCards(ctx context.Context, userUid uuid.UUID) error {
	log.Printf("paymentsRepository.DeleteUserPaymentCards: user uid %s", userUid)

	cardRow := pgEntity.NewUserPaymentCardRow().FromEntity(entity.UserPaymentCard{UserUid: userUid})

	if err := r.Delete(ctx, cardRow, cardRow.ConditionUserUidEqual()); err != nil {
		log.Printf("failed to delete payment cards by user uid %s: %v", userUid, err)
		return errors.WithStack(err)
	}
	return nil
}
