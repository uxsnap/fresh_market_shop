package httpEntity

import uuid "github.com/satori/go.uuid"

type AuthUserInfo struct {
	UserUid     uuid.UUID
	Role        string
	Permissions string
}
