package router

import (
	v1 "gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func  InitMessageRouter(Router *gin.RouterGroup){
	MessageRouter := Router.Group("message")
	{
		MessageRouter.POST("createMessage", v1.CreateMessage)   // 新建Question
		MessageRouter.GET("findPaper", v1.FindPaper)   // 新建Question
	}
}