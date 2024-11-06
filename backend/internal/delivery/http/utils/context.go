package httpUtils

import (
	"context"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
)

func GetUserInfoFromContext(ctx context.Context) httpEntity.AuthUserInfo {
	info, ok := ctx.Value(httpEntity.AuthUserInfo{}).(httpEntity.AuthUserInfo)

	if !ok {
		return httpEntity.AuthUserInfo{}
	}

	return info
}
