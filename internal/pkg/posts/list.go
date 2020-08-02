package posts

import (
	"context"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/dbcon"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/models"
)

type Posts struct {
	db *dbcon.Db
}

func New(db *dbcon.Db) *Posts {
	return &Posts{
		db: db,
	}
}

func (p *Posts) List(ctx context.Context) ([]models.Post, error) {
	const query = "SELECT id, score, views, title, url, upvote, (extract(epoch from created) :: bigint) as created FROM posts"
	posts := []models.Post{}
	err := p.db.Sqlx.SelectContext(ctx, &posts, query)
	return posts, err
}
