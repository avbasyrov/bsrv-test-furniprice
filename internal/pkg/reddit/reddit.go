package reddit

import "github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/interfaces"

type Reddit struct {
	posts    interfaces.PostRepository
	comments interfaces.CommentsRepository
}

func New(posts interfaces.PostRepository,
	comments interfaces.CommentsRepository,
) *Reddit {
	return &Reddit{
		posts:    posts,
		comments: comments,
	}
}
