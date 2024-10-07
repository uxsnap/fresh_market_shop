package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// пока как в сервисе авторизации
type AuthUser struct {
	Uid         uuid.UUID
	Email       string
	Role        UserRole
	Permissions []UserPermission
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (u *AuthUser) PermissionsStrings() []string {
	res := make([]string, len(u.Permissions))
	for i := 0; i < len(res); i++ {
		res[i] = string(u.Permissions[i])
	}
	return res
}

func PermissionsFromStrings(ps []string) []UserPermission {
	perms := make([]UserPermission, len(ps))
	for i := 0; i < len(ps); i++ {
		perms[i] = UserPermission(ps[i])
	}
	return perms
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
