package auth

import "github.com/acgtubio/eydsh-auth/model"

type AuthService interface {
	Register(model.User) error
	Authenticate(model.User) (bool, error)
}
