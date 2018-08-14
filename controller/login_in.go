package controller

import (
	"github.com/gin-gonic/gin"
	"smile-by/utils"
	"net/http"
	"log"
	"smile-by/model"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func Login_in() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//用户登陆
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
		user := model.User{}
		err_qu := coll.FindId(bson.ObjectIdHex(uid)).One(&user)

		if len(user.Id_) ==0 || err_qu != nil {
			fmt.Println("查询日志",err_qu)
			ctx.JSON(http.StatusOK, gin.H{"message":"用户未注册"})
			return
		}

		conn := *utils.RedisConn()
		fmt.Println("redis 连接对象2的内存地址",&conn)
		//释放redis资源
		defer conn.Close()
		_,err2 := conn.Do("SET","user_"+user.Id_.Hex(),"{\"uid\":"+uid+",\"name\":\"nil\"}","EX",45*60)
		if err2 != nil {log.Println(err2)}
		ctx.JSON(http.StatusOK,"OK")
	}
}
