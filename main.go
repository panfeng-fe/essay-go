package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/panfeng-fe/essay-go/route"
	// "github.com/kataras/iris/v12/middleware/logger"
)

func main() {
	app := iris.New()

	app.Configure(iris.WithConfiguration(iris.Configuration{
		Tunneling:                         iris.TunnelingConfiguration{},
		IgnoreServerErrors:                nil,
		DisableStartupLog:                 false,
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		DisablePathCorrectionRedirection:  false,
		EnablePathEscape:                  false,
		EnableOptimizations:               false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "2006-01-02 15:04:05",
		Charset:                           "utf-8",
		PostMaxMemory:                     0,
		LocaleContextKey:                  "",
		ViewLayoutContextKey:              "",
		ViewDataContextKey:                "",
		RemoteAddrHeaders:                 nil,
		Other:                             nil,
	}))

	// 目前不添加log
	// app.Logger().SetOutput(io.MultiWriter(logFile, os.Stdout))
	// requestLogger := logger.New(logger.Config{
	// 	Status: true,
	// 	IP:     true,
	// 	Method: true,
	// 	Path:   true,
	// 	Query:  true,
	// })
	// app.Use(requestLogger)

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Authorization"},
		Debug:            false,
		AllowCredentials: true,
	})

	router := app.Party("/", crs).AllowMethods(iris.MethodOptions)

	// 路由注册
	route.RouterInit(router)

	app.Run(
		iris.Addr(":8888"),
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}
