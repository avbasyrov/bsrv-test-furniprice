package reddit

import "context"

func (r *Reddit) UpVote(ctx context.Context, userID int, postID string) error {
	return r.posts.VoteUp(ctx, postID, userID)
}

func (r *Reddit) UnVote(ctx context.Context, userID int, postID string) error {
	return r.posts.UnVote(ctx, postID, userID)
}

func (r *Reddit) DownVote(ctx context.Context, userID int, postID string) error {
	return r.posts.VoteDown(ctx, postID, userID)
}
