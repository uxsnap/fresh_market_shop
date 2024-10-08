package useCaseUsers

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func (uc *UseCaseUsers) DeleteUser(ctx context.Context, uid uuid.UUID) error {
	log.Printf("ucUsers.DeleteUser: uid %s", uid)

	if err := uc.usersRepository.DeleteUserByUid(ctx, uid); err != nil {
		log.Printf("failed to delete user %s: %v", uid, err)
		return errors.WithStack(err)
	}

	return nil
}
