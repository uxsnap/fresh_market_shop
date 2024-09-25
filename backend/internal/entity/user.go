package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// пока как в сервисе авторизации
type User struct {
	Uid         uuid.UUID
	Email       string
	Role        UserRole
	Permissions []UserPermission
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (u *User) PermissionsStrings() []string {
	res := make([]string, len(u.Permissions))
	for i := 0; i < len(res); i++ {
		res[i] = string(u.Permissions[i])
	}
	return res
}

type UserPermission string

const (
	UserPermissionNotVerified = UserPermission("not_verified")
	UserPermissionBase        = UserPermission("base")
)

type UserRole string

const (
	UserRoleUser  = UserRole("user")
	UserRoleAdmin = UserRole("admin")
)