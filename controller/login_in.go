package controller

import (
	"github.com/gin-gonic/gin"
	"smile-by/utils"
	"net/http"
	"log"
	"smile-by/model"
	"gopkg.in/mgo.v2/bson"
)

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
		//TODO mongo 持久化
		coll := utils.ShowAdminDB().C("user")

		objectId := bson.NewObjectId();

		err1 := coll.Insert(&model.User{objectId,"nil",25})
		if err1 !=nil {
			log.Println("持久化异常",err1)
			panic(err1)
		}

		con :=utils.RedisPool.Get()
		//释放redis资源
		defer con.Close()
		_,err2 := con.Do("SET","user_"+objectId.Hex(),"{\"uid\":"+uid+",\"name\":\"nil\"}","EX",45*60)
		if err2 != nil {log.Println(err2)}
		ctx.JSON(http.StatusOK,"OK")
	}
}
