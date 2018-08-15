package controller

import (
	"github.com/gin-gonic/gin"
	"smile-by/utils"
	"net/http"
	"smile-by/model"
	"gopkg.in/mgo.v2/bson"
	"github.com/gin-gonic/gin/json"
	"smile-by/config"
)

func Login_in() gin.HandlerFunc {

	logger := config.Log
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
			logger.Info("查询日志",err_qu)
			ctx.JSON(http.StatusOK, gin.H{"message":"用户未注册"})
			return
		}

		conn := utils.RedisPool.Get()
		//释放redis资源
		defer conn.Close()
		//转化json
		jsons,_ :=json.Marshal(user)
		_,err2 := conn.Do("SET","user_"+user.Id_.Hex(),jsons,"EX",45*60)
		if err2 != nil {logger.Info(err2)}
		ctx.JSON(http.StatusOK,"OK")
	}
}
