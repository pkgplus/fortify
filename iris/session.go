package iris

import (
	"net/http"

	"github.com/kataras/iris/context"
)

const (
	CONTEXT_OPENID_TAG = "OpenID"
	CONTEXT_UNION_TAG  = "UnionID"
)

func SessionCheck(ctx context.Context) {
	// storage
	store, ok := GetStorage(ctx)
	if !ok {
		SendResponse(ctx, http.StatusInternalServerError, "context exception for getting storage", "")
		return
	}

	sid := ctx.Params().Get("sid")
	resp, err := store.QuerySession(sid)
	if err != nil {
		SendResponse(ctx, http.StatusUnauthorized, "session maybe expired", err.Error())
		return
	}

	ctx.Values().Set(CONTEXT_OPENID_TAG, resp.OpenID)
	ctx.Values().Set(CONTEXT_UNION_TAG, resp.Unionid)

	ctx.Next()
}

func GetOpenID(ctx context.Context) string {
	return ctx.Values().GetString(CONTEXT_OPENID_TAG)
}

func GetUnionID(ctx context.Context) string {
	return ctx.Values().GetString(CONTEXT_UNION_TAG)
}
