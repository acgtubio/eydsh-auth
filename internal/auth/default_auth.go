package auth

import (
	"github.com/acgtubio/eydsh-auth/model"
)

type defaultAuth struct{}

func DefaultAuthService() (*defaultAuth, error) {
	return &defaultAuth{}, nil
}

func (*defaultAuth) Register(user model.User) error {
	println("registered.")
	println(user.Username)

	return nil
}

func (*defaultAuth) Authenticate(user model.User) (bool, error) {
	println("authenticated.")
	println(user.Username)

	return true, nil
}
