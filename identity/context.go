package identity

import "context"

type theKey struct{}

func ContextWithIdentity(ctx context.Context, id Identity) context.Context {
	return context.WithValue(ctx, theKey{}, id)
}

func IdentityFromContext(ctx context.Context) Identity {
	u, ok := ctx.Value(theKey{}).(Identity)
	if ok && u != nil {
		return u
	}
	return nil
}
