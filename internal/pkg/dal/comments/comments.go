package comments

import (
	"context"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/dbcon"
)

type Comments struct {
	db *dbcon.Db
}

func New(db *dbcon.Db) *Comments {
	return &Comments{
		db: db,
	}
}

func (c *Comments) Add(ctx context.Context, postID string, comment string, authorID int) error {
	const query = "INSERT INTO public.comments (post_id, body, author_id) VALUES ($1, $2, $3)"

	_, err := c.db.Sqlx.ExecContext(ctx, query, postID, comment, authorID)
	return err
}

func (c *Comments) Delete(ctx context.Context, commentID string, authorID int) error {
	const query = "DELETE FROM public.comments WHERE id = $1 AND author_id = $2"

	_, err := c.db.Sqlx.ExecContext(ctx, query, commentID, authorID)
	return err
}
