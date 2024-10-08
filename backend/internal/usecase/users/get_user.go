package useCaseUsers

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseUsers) GetUser(ctx context.Context, uid uuid.UUID) (entity.User, bool, error) {
	log.Printf("ucUsers.GetUser: uid %s", uid)

	user, isFound, err := uc.usersRepository.GetUserByUid(ctx, uid)
	if err != nil {
		log.Printf("failed to get user %s: %v", uid, err)
		return entity.User{}, false, errors.WithStack(err)
	}

	return user, isFound, nil
}
