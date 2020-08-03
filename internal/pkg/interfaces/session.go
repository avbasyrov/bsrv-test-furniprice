package interfaces

import "context"

type Session interface {
	Write(context.Context, int) (int64, error)
	Load(context.Context, int64) (int, error)
}
