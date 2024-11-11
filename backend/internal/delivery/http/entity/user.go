package httpEntity

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type User struct {
	Uid       uuid.UUID `json:"uid"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Birthday  time.Time `json:"birthday"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func UserFromEntity(user entity.User) User {
	return User{
		Uid:       user.Uid,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Birthday:  user.Birthday,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UserToEntity(user User) entity.User {
	return entity.User{
		Uid:       user.Uid,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Birthday:  user.Birthday,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
