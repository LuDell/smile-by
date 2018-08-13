package controller

import (
	"github.com/gin-gonic/gin"
	"smile-by/utils"
	"gopkg.in/mgo.v2/bson"
	"smile-by/model"
	"log"
)

func Register() gin.HandlerFunc {
	return func(ctx *gin.Context){
		//用户注册
		uid := ctx.Query("uid")
		//TODO mongo 持久化
		coll := utils.ShowAdminDB().C("user")

		objectId := bson.NewObjectId();

		err1 := coll.Insert(&model.User{objectId,uid,25})
		if err1 !=nil {
			log.Println("持久化异常",err1)
			panic(err1)
		}
	}
}
