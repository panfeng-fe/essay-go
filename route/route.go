package route

import (
	"github.com/kataras/iris/v12"
)

func RouterInit(router iris.Party) {
	EssayInit(router.Party("/essay"))
}
