package handlers

import (
	"net/http"

	"github.com/kataras/iris/context"
	"github.com/xuebing1110/fortify/iris"
	"github.com/xuebing1110/fortify/pkg/wechat"
)

type LoginReq struct {
	Code string `json:"code"`
}

type LoginResp struct {
	*iris.Response
	Session string `json:"session"`
}

func UserLogin(ctx context.Context) {
	lr := new(LoginReq)

	// request
	err := ctx.ReadJSON(lr)
	if err != nil {
		iris.SendResponse(ctx, http.StatusBadRequest, "Parse to json failed", err.Error())
		return
	}
	if lr.Code == "" {
		iris.SendResponse(ctx, http.StatusBadRequest, "code is required", "")
		return
	}

	// storage
	store, ok := iris.GetStorage(ctx)
	if !ok {
		iris.SendResponse(ctx, http.StatusInternalServerError, "context exception for getting storage", "")
		return
	}

	// openid
	sessRet, err := wechat.Jscode2Session(lr.Code)
	if err != nil {
		iris.SendResponse(ctx, http.StatusInternalServerError, "jscode2session failed", err.Error())
		return
	}

	// create session
	sess_3rd := sessRet.OpenID
	// sess_3rd, err := user.GetRandomID(16)
	// if err != nil {
	//  iris.SendResponse(ctx, http.StatusInternalServerError, "create 3rd_sess failed", err.Error())
	//  return
	// }

	// storage
	err = store.SaveSession(sess_3rd, sessRet)
	if err != nil {
		iris.SendResponse(ctx, http.StatusInternalServerError, "save sess_3rd and sessinfo failed", err.Error())
		return
	}

	ctx.JSON(&LoginResp{Session: sess_3rd})
}
