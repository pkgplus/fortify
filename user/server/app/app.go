package app

import (
	"github.com/kataras/iris"
	irircom "github.com/xuebing1110/fortify/iris"
)

var (
	IrisApp *iris.Application
)

func init() {
	IrisApp = iris.New()

	// storage
	IrisApp.Use(irircom.UseRedisStorage)
}

func GetIrisApp() *iris.Application {
	return IrisApp
}
