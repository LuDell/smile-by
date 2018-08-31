package controller

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"smile-by/model"
	"github.com/jeanphorn/log4go"
	"smile-by/utils"
)

func Register() gin.HandlerFunc {

	logger := log4go.NewDefaultLogger(log4go.INFO)
	return func(ctx *gin.Context){
		//用户注册
		uid := ctx.Query("uid")
		//TODO mongo 持久化
		DB := utils.ShowDB()
		defer DB.Session.Close()
		coll := DB.C("user")

		objectId := bson.NewObjectId()

		err1 := coll.Insert(&model.User{objectId,uid,25})
		if err1 !=nil {
			logger.Info("持久化异常",err1)
		}
	}
}
