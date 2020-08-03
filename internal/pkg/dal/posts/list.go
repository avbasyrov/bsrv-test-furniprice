package posts

import (
	"context"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/dbcon"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/interfaces"
)

type Posts struct {
	db *dbcon.Db
}

func New(db *dbcon.Db) *Posts {
	return &Posts{
		db: db,
	}
}

const selectQuery = "SELECT p.id, p.views, p.title, p.url, p.text, p.created, " +
	"(SELECT COALESCE(SUM(v.vote), 0) AS score FROM public.votes v WHERE v.post_id = p.id) AS score, " +
	"(SELECT CASE WHEN COUNT(vv.vote) > 0 THEN 100 * (SUM(vv.vote) + COUNT(vv.vote)) / (2 * COUNT(vv.vote)) ELSE 100 END AS UpvotePercentage FROM public.votes vv WHERE vv.post_id = p.id) AS UpvotePercentage, " +
	"p.category, p.author_id AS AuthorID, u.login AS AuthorName, " +
	"CASE WHEN p.is_link IS TRUE THEN 'link' ELSE 'text' END AS type,  " +
	"(SELECT COALESCE(json_agg(json_build_object('user', v.user_id, 'vote', v.vote)), '[]') FROM votes v WHERE v.post_id = p.id) AS votes " +
	"FROM posts p " +
	"LEFT JOIN users u ON u.id = p.author_id "

func (p *Posts) List(ctx context.Context) ([]interfaces.Post, error) {
	var posts []interfaces.Post
	err := p.db.Sqlx.SelectContext(ctx, &posts, selectQuery)
	return posts, err
}

func (p *Posts) ByCategory(ctx context.Context, category string) ([]interfaces.Post, error) {
	const query = selectQuery + "WHERE p.category = $1"
	var posts []interfaces.Post
	err := p.db.Sqlx.SelectContext(ctx, &posts, query, category)
	return posts, err
}

func (p *Posts) GetByID(ctx context.Context, postID string) (interfaces.Post, error) {
	const query = selectQuery + "WHERE p.id = $1"
	var post interfaces.Post
	err := p.db.Sqlx.GetContext(ctx, &post, query, postID)
	return post, err
}

func (p *Posts) ByAuthor(ctx context.Context, authorID int) ([]interfaces.Post, error) {
	const query = selectQuery + "WHERE p.author_id = $1"
	var posts []interfaces.Post
	err := p.db.Sqlx.SelectContext(ctx, &posts, query, authorID)
	return posts, err
}

func (p *Posts) Delete(ctx context.Context, postID string) error {
	_, err := p.db.Sqlx.ExecContext(ctx, "DELETE FROM public.posts WHERE id = $1", postID)
	return err
}
