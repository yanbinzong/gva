package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitQuestion_optionsRouter(Router *gin.RouterGroup) {
	Question_optionsRouter := Router.Group("question_options").Use(middleware.OperationRecord())
	{
		Question_optionsRouter.POST("createQuestion_options", v1.CreateQuestion_options)   // 新建Question_options
		Question_optionsRouter.DELETE("deleteQuestion_options", v1.DeleteQuestion_options) // 删除Question_options
		Question_optionsRouter.DELETE("deleteQuestion_optionsByIds", v1.DeleteQuestion_optionsByIds) // 批量删除Question_options
		Question_optionsRouter.PUT("updateQuestion_options", v1.UpdateQuestion_options)    // 更新Question_options
		Question_optionsRouter.GET("findQuestion_options", v1.FindQuestion_options)        // 根据ID获取Question_options
		Question_optionsRouter.GET("getQuestion_optionsList", v1.GetQuestion_optionsList)  // 获取Question_options列表
	}
}
