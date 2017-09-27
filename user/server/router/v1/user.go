package v1

import (
	"github.com/xuebing1110/fortify/user/server/app"
	"github.com/xuebing1110/fortify/user/server/handlers"
)

func init() {
	irisApp := app.GetIrisApp()

	api := irisApp.Party("/api/v2/fortify")

	// user
	api.Post("/users", handlers.UserRegiste)

	// session
	api.Post("/login", handlers.UserLogin)
}
