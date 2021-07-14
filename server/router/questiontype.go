package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitQuestion_typeRouter(Router *gin.RouterGroup) {
	Question_typeRouter := Router.Group("question_type").Use(middleware.OperationRecord())
	{
		Question_typeRouter.POST("createQuestion_type", v1.CreateQuestion_type)   // 新建Question_type
		Question_typeRouter.DELETE("deleteQuestion_type", v1.DeleteQuestion_type) // 删除Question_type
		Question_typeRouter.DELETE("deleteQuestion_typeByIds", v1.DeleteQuestion_typeByIds) // 批量删除Question_type
		Question_typeRouter.PUT("updateQuestion_type", v1.UpdateQuestion_type)    // 更新Question_type
		Question_typeRouter.GET("findQuestion_type", v1.FindQuestion_type)        // 根据ID获取Question_type
		Question_typeRouter.GET("getQuestion_typeList", v1.GetQuestion_typeList)  // 获取Question_type列表
		Question_typeRouter.GET("getQuestion_typeList_all", v1.GetQuestion_typeList_all) //获取Question_typeList_all全部列表
	}
}
