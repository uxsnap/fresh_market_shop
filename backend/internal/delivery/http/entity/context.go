package httpEntity

import (
	"context"
	"errors"
)

type ctxKeyAuthUserInfo struct{}

func ContextWithAuthUserInfo(ctx context.Context, info AuthUserInfo) context.Context {
	return context.WithValue(ctx, ctxKeyAuthUserInfo{}, info)
}

func AuthUserInfoFromContext(ctx context.Context) (AuthUserInfo, error) {
	userInfoFromCtx := ctx.Value(ctxKeyAuthUserInfo{})
	userInfo, ok := userInfoFromCtx.(AuthUserInfo)
	if !ok {
		return AuthUserInfo{}, errors.New("failed to get user from ctx")
	}

	return userInfo, nil
}
