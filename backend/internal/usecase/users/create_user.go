package useCaseUsers

import (
	"context"
	"log"
	"net/mail"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (uc *UseCaseUsers) CreateUser(ctx context.Context, user entity.User) (uuid.UUID, *errorWrapper.Error) {
	log.Printf("ucUsers.CreateUser: email %s username %s", user.Email, user.FirstName)

	if _, err := mail.ParseAddress(user.Email); err != nil {
		return uuid.UUID{}, errorWrapper.NewError(
			errorWrapper.UserEmailError, "неправильный формат email",
		)
	}

	if uuid.Equal(user.Uid, uuid.UUID{}) {
		user.Uid = uuid.NewV4()
	}

	user.CreatedAt = time.Now().UTC()

	if err := uc.usersRepository.CreateUser(ctx, user); err != nil {
		log.Printf("failed to create user:%v", err)
		return uuid.UUID{}, errorWrapper.NewError(
			errorWrapper.UserInfoError, "не удалось создать пользователя",
		)
	}
	return user.Uid, nil
}
