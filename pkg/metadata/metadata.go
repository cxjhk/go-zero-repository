package metadata

import "context"

func WithMetadata(ctx context.Context, key, val any) context.Context {
	return context.WithValue(ctx, key, val)
}

func GetMetadataFromCtx(ctx context.Context, key any) any {
	return ctx.Value(key)
}
