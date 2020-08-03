package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/interfaces"
	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	jwtSecret      []byte
	users          interfaces.UsersRepository
	sessionManager interfaces.Session
}

func New(jwtSecret []byte, sessionManager interfaces.Session, users interfaces.UsersRepository) *Auth {
	return &Auth{
		jwtSecret:      jwtSecret,
		users:          users,
		sessionManager: sessionManager,
	}
}

func (a *Auth) BeginSession(ctx context.Context, userID int, userLogin string) (string, error) {
	sessionID, err := a.sessionManager.Write(ctx, userID)
	if err != nil {
		return "", errors.New("unable to save session data")
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": jwt.MapClaims{
			"username": userLogin,
			"id":       userID,
		},
		"sessionID": sessionID,
	}).SignedString(a.jwtSecret)
}

func (a *Auth) GetAuthorized(ctx context.Context, jwtToken string) (int, bool, error) {
	if len(jwtToken) > 7 {
		// remove "Bearer " if exists
		jwtToken = jwtToken[7:]
	}
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return a.jwtSecret, nil
	})
	if err != nil {
		return 0, false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, false, err
	}

	sessionID, ok := claims["sessionID"].(float64)
	if !ok {
		return 0, false, errors.New("broken sessionID")
	}

	userID, err := a.sessionManager.Load(ctx, int64(sessionID))
	if err != nil {
		return 0, false, err
	}

	user, err := a.users.GetByID(ctx, userID)
	if err != nil {
		return 0, false, err
	}

	return user.ID, user.Role == "admin", err
}
