package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"smile-by/handler"
	"smile-by/controller"
	"smile-by/resource"
	"smile-by/utils"
)

func main()  {
	//初始init
	logger := utils.Logger
	dirs := []string{"views","config"} // 设置需要释放的目录
	for _, dir := range dirs {
		// 解压dir目录到当前目录
		if err := resource.RestoreAssets("./", dir); err != nil {
			break
		}
	}

	app := gin.New()
	app.Use(gin.Recovery())
	app.Static("/static","views/static")
	app.StaticFile("/favicon.ico","views/static/favicon.ico")
	app.LoadHTMLGlob("views/*.html")

	app.GET("/",func(ctx *gin.Context) {

		cookie,err := ctx.Request.Cookie("admin")
		var value *string
		if(err == nil){
			value = &cookie.Value
		}

		logger.Info("返回的cookie值", value)
		ctx.HTML(http.StatusOK,"index.html",gin.H{"data":"hello world"})

	})

	group := app.Group("api",handler.AuthHandler)
	{
		group.OPTIONS("public_info", func(ctx *gin.Context) {

			val,bol := ctx.Get("landing")

			ctx.JSON(http.StatusOK,gin.H{"val":val,"bol":bol,"data":"hello world"})
		})

	}

	app.GET("login_in",controller.Login_in)

	app.GET("register",controller.Register)

	server := &http.Server{
		Addr: ":8088",
		Handler: app,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 20*time.Second,
		MaxHeaderBytes: 1<<20,
	}
	server.ListenAndServe()
}