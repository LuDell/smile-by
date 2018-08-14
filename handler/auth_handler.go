package handler

import (
	"github.com/gin-gonic/gin"
	"smile-by/utils"
)

func AuthHandler() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		conn := *utils.RedisConn()
		defer conn.Close()
		val,err := conn.Do("GET","user_"+token)
		if err != nil || val == nil {
			ctx.Abort()
		}
		ctx.Set("landing",true)

		ctx.Next()
	}
}
