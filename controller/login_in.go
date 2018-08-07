package controller

import (
	"github.com/gin-gonic/gin"
	"smile-by/utils"
	"net/http"
	"log"
	"fmt"
)

type Person struct {
	Uid string
	Name string
}

func Login_in() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		uid_cookie:=&http.Cookie{
			Name:   "admin",
			Value:   "123456",
			Path:     "/",
			HttpOnly: true,
			MaxAge:   -1,
		}
		http.SetCookie(ctx.Writer,uid_cookie)

		uid := ctx.Query("uid")
		guid:= utils.UniqueId()
		//TODO mongo 持久化
		coll := utils.Db.C("user")
		err1 := coll.Insert(&Person{uid,"nil"})
		if err1 !=nil {
			fmt.Println("mongo持久化异常",err1)
		}

		con :=utils.RedisPool.Get()
		//释放redis资源
		defer con.Close()
		_,err2 := con.Do("SET","user_"+guid,"{\"uid\":"+uid+",\"name\":\"nil\"}","EX",45*60)
		if err2 != nil {log.Println(err2)}
		ctx.JSON(http.StatusOK,"OK")
	}
}
