package session

import (
	"context"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/dbcon"
)

type Session struct {
	db *dbcon.Db
}

func New(db *dbcon.Db) *Session {
	return &Session{
		db: db,
	}
}

func (s *Session) Write(ctx context.Context, userID int) (int64, error) {
	const query = "INSERT INTO public.sessions (user_id) VALUES ($1) RETURNING id"

	sessionID := int64(0)
	row := s.db.Sqlx.QueryRowContext(ctx, query, userID)
	err := row.Scan(&sessionID)

	return sessionID, err
}

func (s *Session) Load(ctx context.Context, sessionID int64) (int, error) {
	const query = "SELECT user_id FROM public.sessions WHERE id = $1"
	userID := 0
	err := s.db.Sqlx.GetContext(ctx, &userID, query, sessionID)
	return userID, err
}
