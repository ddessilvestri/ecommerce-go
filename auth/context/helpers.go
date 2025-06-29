package context

import (
	"context"

	"github.com/ddessilvestri/ecommerce-go/models"
)

func WithUser(ctx context.Context, u *models.AuthUser) context.Context {
	return context.WithValue(ctx, AuthUserKey(), u)
}

func UserFromContext(ctx context.Context) (*models.AuthUser, bool) {
	u, ok := ctx.Value(AuthUserKey()).(*models.AuthUser)
	return u, ok
}

func IsAdmin(ctx context.Context) bool {
	_, ok := UserFromContext(ctx)
	return ok && true
}
