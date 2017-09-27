package handlers

import (
	"net/http"

	"github.com/kataras/iris/context"
	"github.com/xuebing1110/fortify/iris"
	"github.com/xuebing1110/fortify/user/models"
)

func UserRegiste(ctx context.Context) {
	user := new(models.User)

	// request
	err := ctx.ReadJSON(user)
	if err != nil {
		iris.SendResponse(ctx, http.StatusBadRequest, "Parse json to User failed", err.Error())
		return
	}

	// storage
	store, ok := iris.GetStorage(ctx)
	if !ok {
		iris.SendResponse(ctx, http.StatusInternalServerError, "context exception for getting storage", "")
		return
	}

	// save
	err = store.AddUser(*user)
	if err != nil {
		iris.SendResponse(ctx, http.StatusInternalServerError, "save user failed", err.Error())
		return
	}

	iris.SendResponse(ctx, http.StatusCreated, "OK", "")
}
