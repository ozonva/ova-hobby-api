package repo

import (
	"context"
	"google.golang.org/grpc"
)

var repoKey = "repo"

// NewContext ...
func NewContext(ctx context.Context, repo Repo) context.Context {
	ctxDB := context.WithValue(ctx, &repoKey, repo)
	return ctxDB
}

// FromContext ...
func FromContext(ctx context.Context) Repo {
	client, ok := ctx.Value(&repoKey).(Repo)
	if !ok {
		panic("Error getting connection from context")
	}
	return client
}

// NewInterceptorWithRepo ...
func NewInterceptorWithRepo(repo Repo) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		return handler(NewContext(ctx, repo), req)
	}
}
