package useCaseUsers

import (
	"context"
	"log"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (uc *UseCaseUsers) GetUser(ctx context.Context, uid uuid.UUID) (entity.User, bool, *errorWrapper.Error) {
	log.Printf("ucUsers.GetUser: uid %s", uid)

	user, isFound, err := uc.usersRepository.GetUserByUid(ctx, uid)
	if err != nil {
		log.Printf("failed to get user %s: %v", uid, err)
		return entity.User{}, false, errorWrapper.NewError(errorWrapper.UserInfoError, "пользователь не был найден")
	}

	return user, isFound, nil
}
