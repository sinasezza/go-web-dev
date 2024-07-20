package context

import (
	"context"

	"github.com/sinasezza/go-web-dev/models"
)

type key string

const (
	userKey key = "user"
)

func WithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func User(ctx context.Context) *models.User {
	val := ctx.Value(userKey)
	user, ok := val.(*models.User)
	if !ok {
		return nil
	}
	return user
}

func IsAuthenticated(ctx context.Context) bool {
	return ctx.Value(userKey) != nil
}
