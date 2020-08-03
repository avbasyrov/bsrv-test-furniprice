package reddit

import "github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/interfaces"

type Reddit struct {
	auth     interfaces.AuthManager
	users    interfaces.UsersRepository
	posts    interfaces.PostRepository
	comments interfaces.CommentsRepository
}

func New(auth interfaces.AuthManager,
	users interfaces.UsersRepository,
	posts interfaces.PostRepository,
	comments interfaces.CommentsRepository,
) *Reddit {
	return &Reddit{
		auth:     auth,
		users:    users,
		posts:    posts,
		comments: comments,
	}
}
