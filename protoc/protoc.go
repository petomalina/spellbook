package protoc

import "context"

type Plugin interface {
	Build(ctx context.Context) error
}
