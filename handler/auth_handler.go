package handler

import (
	"github.com/gin-gonic/gin"
	"smile-by/utils"
)

func AuthHandler() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		con := utils.RedisPool.Get()
		val,err := con.Do("GET","user_"+token)
		if err != nil || val == nil {
			ctx.Abort()
		}
		ctx.Set("landing",true)

		ctx.Next()
	}
}
