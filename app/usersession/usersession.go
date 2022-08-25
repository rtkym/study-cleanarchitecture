package usersession

import "context"

type UserSession struct {
	UserID string
}

type key int

const ctxKey key = iota

func NewContext(ctx context.Context, userSession *UserSession) context.Context {
	return context.WithValue(ctx, ctxKey, userSession)
}

func FromContext(ctx context.Context) (*UserSession, bool) {
	v, ok := ctx.Value(ctxKey).(*UserSession)
	return v, ok
}
