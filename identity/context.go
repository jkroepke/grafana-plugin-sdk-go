package identity

import "context"

type theUserKey struct{}

func ContextWithIdentity(ctx context.Context, sub Identity) context.Context {
	return context.WithValue(ctx, theUserKey{}, sub)
}

func IdentityFromContext(ctx context.Context) Identity {
	u, ok := ctx.Value(theUserKey{}).(Identity)
	if ok && u != nil {
		return u
	}
	return nil
}
