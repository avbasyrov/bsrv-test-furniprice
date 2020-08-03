package interfaces

import "context"

type AuthManager interface {
	BeginSession(ctx context.Context, userID int, userLogin string) (jwtToken string, err error)
	GetAuthorized(ctx context.Context, jwtToken string) (userID int, isAdmin bool, err error)
}
