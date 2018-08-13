package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io"
	"os"
	"time"
	"smile-by/handler"
	"smile-by/controller"
	"smile-by/utils"
)

func main()  {
	//记录日志到日志文件
	gin.DefaultWriter = io.MultiWriter(utils.LogFile(),os.Stdout)

	app := gin.Default()

	app.LoadHTMLGlob("views/*")

	app.GET("/",func(ctx *gin.Context) {

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