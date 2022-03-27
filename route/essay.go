package route

import (
	"github.com/kataras/iris/v12"
)

func EssayInit(router iris.Party) {
	appInit(router.Party("/app"))
}

func appInit(router iris.Party) {
	router.Get("/test", Test)
}

func Test(ctx iris.Context) {
	ctx.JSON("111")
}
