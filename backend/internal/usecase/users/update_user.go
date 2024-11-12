package useCaseUsers

import (
	"context"
	"fmt"
	"log"
	"net/mail"
	"time"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (uc *UseCaseUsers) UpdateUser(ctx context.Context, user entity.User) *errorWrapper.Error {
	log.Printf("ucUsers.UpdateUser: email %s firstName %s lastName %s", user.Email, user.FirstName, user.LastName)

	savedUser, isFound, err := uc.usersRepository.GetUserByUid(ctx, user.Uid)
	if err != nil {
		log.Printf("failed to update user %s: %v", user.Uid, err)
		return errorWrapper.NewError(
			errorWrapper.UserInfoError, "не удалось найти пользователя с таким uid",
		)
	}
	if !isFound {
		log.Printf("failed to update user %s: user not found", user.Uid)
		return errorWrapper.NewError(
			errorWrapper.UserInfoError, "не найден пользователь с таким uid",
		)
	}

	if _, err = mail.ParseAddress(user.Email); err != nil {
		return errorWrapper.NewError(
			errorWrapper.UserEmailError, "неправильный формат email",
		)
	}

	if len(user.FirstName) < 2 {
		return errorWrapper.NewError(
			errorWrapper.UserNameError, "длина имени и фамилии пользователя должна быть больше 1",
		)
	}

	fmt.Println(time.Since(savedUser.UpdatedAt).Minutes())

	if time.Since(savedUser.UpdatedAt).Minutes() < 15 {
		return errorWrapper.NewError(
			errorWrapper.UserInfoError, "данные пользователя можно обновлять раз в 15 минут",
		)
	}

	if err := uc.usersRepository.UpdateUser(ctx, user); err != nil {
		log.Printf("failed to update user %s: %v", user.Uid, err)
		return errorWrapper.NewError(
			errorWrapper.UserInfoError, "не получилось обновить пользователя",
		)
	}
	return nil
}
