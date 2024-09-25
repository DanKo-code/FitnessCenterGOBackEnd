package auth

import "context"

type UseCase interface {
	SignUp(ctx context.Context, FirstName string, LastName string, Email string, password string) error
}
