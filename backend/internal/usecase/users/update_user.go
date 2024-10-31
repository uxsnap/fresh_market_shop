package useCaseUsers

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseUsers) UpdateUser(ctx context.Context, user entity.User) error {
	log.Printf("ucUsers.UpdateUser: email %s username %s", user.Email, user.Username)

	savedUser, isFound, err := uc.usersRepository.GetUserByUid(ctx, user.Uid)
	if err != nil {
		log.Printf("failed to update user %s: %v", user.Uid, err)
		return errors.WithStack(err)
	}
	if !isFound {
		log.Printf("failed to update user %s: user not found", user.Uid)
		return errors.New("user not found")
	}

	// TODO: несогласованность данных. Сделать отдельный метод для обновления email. а здесь запретить это делать!
	if len(user.Email) != 0 && user.Email != savedUser.Email {
		savedUser.Email = user.Email
	}
	if len(user.Username) != 0 && user.Username != savedUser.Username {
		savedUser.Username = user.Username
	}

	savedUser.UpdatedAt = time.Now().UTC()

	if err := uc.usersRepository.UpdateUser(ctx, savedUser); err != nil {
		log.Printf("failed to update user %s: %v", user.Uid, err)
		return errors.WithStack(err)
	}
	return nil
}
