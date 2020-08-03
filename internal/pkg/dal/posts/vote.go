package posts

import (
	"context"
)

func (p *Posts) VoteUp(ctx context.Context, postID string, userID int) error {
	const query = "INSERT INTO public.votes (post_id, user_id, vote) VALUES ($1, $2, $3) " +
		"ON CONFLICT (post_id, user_id) DO UPDATE SET vote = excluded.vote"

	_, err := p.db.Sqlx.ExecContext(ctx, query, postID, userID, 1)

	return err
}

func (p *Posts) UnVote(ctx context.Context, postID string, userID int) error {
	const query = "DELETE FROM public.votes WHERE post_id = $1 AND user_id = $2"

	_, err := p.db.Sqlx.ExecContext(ctx, query, postID, userID)

	return err
}

func (p *Posts) VoteDown(ctx context.Context, postID string, userID int) error {
	const query = "INSERT INTO public.votes (post_id, user_id, vote) VALUES ($1, $2, $3) " +
		"ON CONFLICT (post_id, user_id) DO UPDATE SET vote = excluded.vote"

	_, err := p.db.Sqlx.ExecContext(ctx, query, postID, userID, -1)

	return err
}
