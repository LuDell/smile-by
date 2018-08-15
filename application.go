package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"smile-by/handler"
	"smile-by/controller"
	"smile-by/config"
)

func main()  {
	//全局日志
	logger := config.Log

	app := gin.New()
	app.Use(gin.Recovery())
	app.Static("/static","views/static")
	app.StaticFile("/favicon.ico","views/static/favicon.ico")
	app.LoadHTMLGlob("views/*.html")

	app.GET("/",func(ctx *gin.Context) {
		logger.Info("gin初始化路由，作者","okay")
		ctx.HTML(http.StatusOK,"index.html",gin.H{"data":"hello world"})

	})

	group := app.Group("api",handler.AuthHandler())
	{
		group.OPTIONS("public_info", func(ctx *gin.Context) {

			val,bol := ctx.Get("landing")

			ctx.JSON(http.StatusOK,gin.H{"val":val,"bol":bol,"data":"hello world"})
		})

	}

	app.GET("login_in",controller.Login_in())

	app.GET("register",controller.Register())

	server := &http.Server{
		Addr: ":8088",
		Handler: app,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 20*time.Second,
		MaxHeaderBytes: 1<<20,
	}
	server.ListenAndServe()
}