package useCaseUsers

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseUsers) CreateUser(ctx context.Context, user entity.User) (uuid.UUID, error) {
	log.Printf("ucUsers.CreateUser: email %s username %s", user.Email, user.Username)

	if err := validateUser(user); err != nil {
		log.Printf("failed to create user: %v", err)
		return uuid.UUID{}, err
	}
	if uuid.Equal(user.Uid, uuid.UUID{}) {
		user.Uid = uuid.NewV4()
	}

	user.CreatedAt = time.Now().UTC()

	if err := uc.usersRepository.CreateUser(ctx, user); err != nil {
		log.Printf("failed to create user:%v", err)
		return uuid.UUID{}, errors.WithStack(err)
	}
	return user.Uid, nil
}
