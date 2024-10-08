package repositoryUsers

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *UsersRepository) CreateUser(ctx context.Context, user entity.User) error {
	log.Printf("usersRepository.CreateUser: uid %s", user.Uid)

	if err := r.Create(ctx, pgEntity.NewUserRow().FromEntity(user)); err != nil {
		log.Printf("failed to create user %s: %v", user.Uid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *UsersRepository) UpdateUser(ctx context.Context, user entity.User) error {
	log.Printf("usersRepository.UpdateUser: uid %s", user.Uid)

	row := pgEntity.NewUserRow().FromEntity(user)

	if err := r.Update(ctx, row, row.ConditionUidEqual()); err != nil {
		log.Printf("failed to update user %s: %v", user.Uid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *UsersRepository) GetUserByUid(ctx context.Context, uid uuid.UUID) (entity.User, bool, error) {
	log.Printf("usersRepository.GetUser: uid %s", uid)

	row := pgEntity.NewUserRow().FromEntity(entity.User{Uid: uid})

	if err := r.GetOne(ctx, row, row.ConditionUidEqual()); err != nil {
		log.Printf("failed to get user %s: %v", uid, err)
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, false, nil
		}

		return entity.User{}, false, errors.WithStack(err)
	}

	return row.ToEntity(), true, nil
}

func (r *UsersRepository) DeleteUserByUid(ctx context.Context, uid uuid.UUID) error {
	log.Printf("usersRepository.DeleteUser: uid %s", uid)

	row := pgEntity.NewUserRow().FromEntity(entity.User{Uid: uid})

	if err := r.Delete(ctx, row, row.ConditionUidEqual()); err != nil {
		log.Printf("failed to delete user %s: %v", uid, err)
		return errors.WithStack(err)
	}
	return nil
}
