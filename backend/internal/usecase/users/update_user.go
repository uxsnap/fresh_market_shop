package useCaseUsers

import (
	"context"
	"log"
	"net/mail"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

func (uc *UseCaseUsers) UpdateUser(ctx context.Context, user entity.User) error {
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

	// if time.Since(savedUser.UpdatedAt).Minutes() < 5 {
	// 	return errorWrapper.NewError(
	// 		errorWrapper.UserInfoError, "данные пользователя можно обновлять раз в 5 минут",
	// 	)
	// }

	if len(user.Email) != 0 && user.Email != savedUser.Email {
		if _, err = mail.ParseAddress(user.Email); err != nil {
			return errorWrapper.NewError(
				errorWrapper.UserEmailError, "неправильный формат email",
			)
		}
		savedUser.Email = user.Email
	}

	// TODO: добавить проверку на занятость email

	if len(user.FirstName) != 0 && user.FirstName != savedUser.FirstName {
		if len(user.FirstName) < 2 {
			return errorWrapper.NewError(
				errorWrapper.UserNameError, "длина имени и фамилии пользователя должна быть больше 1",
			)
		}
		savedUser.FirstName = user.FirstName
	}
	if len(user.LastName) != 0 && user.LastName != savedUser.LastName {
		savedUser.LastName = user.LastName
	}
	if user.Birthday.Unix() != savedUser.Birthday.Unix() {
		savedUser.Birthday = user.Birthday
	}

	if err := uc.usersRepository.UpdateUser(ctx, user); err != nil {
		log.Printf("failed to update user %s: %v", user.Uid, err)
		return errorWrapper.NewError(
			errorWrapper.UserInfoError, "не получилось обновить пользователя",
		)
	}
	return nil
}
