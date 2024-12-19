package authSubrouter

import (
	"log"
	"strings"

	errorWrapper "github.com/uxsnap/fresh_market_shop/backend/internal/error_wrapper"
)

type FullName struct {
	firstName string
	lastName  string
}

func GetUserName(name string) (FullName, *errorWrapper.Error) {
	nameSlice := strings.Split(name, " ")
	firstName := nameSlice[0]
	lastName := ""
	if len(nameSlice) == 2 {
		lastName = nameSlice[1]
	}

	if len(firstName) < 2 {
		log.Printf("failed to validate register user")
		return FullName{}, errorWrapper.NewError(
			errorWrapper.UserNameError, "длина имени пользователя должна быть больше 1",
		)
	}

	return FullName{
		firstName: firstName,
		lastName:  lastName,
	}, nil
}
