package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAnswer_questionRouter(Router *gin.RouterGroup) {
	Answer_questionRouter := Router.Group("answer_question").Use(middleware.OperationRecord())
	{
		Answer_questionRouter.POST("createAnswer_question", v1.CreateAnswer_question)   // 新建Answer_question
		Answer_questionRouter.DELETE("deleteAnswer_question", v1.DeleteAnswer_question) // 删除Answer_question
		Answer_questionRouter.DELETE("deleteAnswer_questionByIds", v1.DeleteAnswer_questionByIds) // 批量删除Answer_question
		Answer_questionRouter.PUT("updateAnswer_question", v1.UpdateAnswer_question)    // 更新Answer_question
		Answer_questionRouter.GET("findAnswer_question", v1.FindAnswer_question)        // 根据ID获取Answer_question
		Answer_questionRouter.GET("getAnswer_questionList", v1.GetAnswer_questionList)  // 获取Answer_question列表
	}
}
