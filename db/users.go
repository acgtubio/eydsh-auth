package db

type Users[T any] interface {
	GetUser() T
	CreateUser(T)
}
