package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitUser_paper_answerRouter(Router *gin.RouterGroup) {
	User_paper_answerRouter := Router.Group("user_paper_answer").Use(middleware.OperationRecord())
	{
		User_paper_answerRouter.POST("createUser_paper_answer", v1.CreateUser_paper_answer)   // 新建User_paper_answer
		User_paper_answerRouter.DELETE("deleteUser_paper_answer", v1.DeleteUser_paper_answer) // 删除User_paper_answer
		User_paper_answerRouter.DELETE("deleteUser_paper_answerByIds", v1.DeleteUser_paper_answerByIds) // 批量删除User_paper_answer
		User_paper_answerRouter.PUT("updateUser_paper_answer", v1.UpdateUser_paper_answer)    // 更新User_paper_answer
		User_paper_answerRouter.GET("findUser_paper_answer", v1.FindUser_paper_answer)        // 根据ID获取User_paper_answer
		User_paper_answerRouter.GET("getUser_paper_answerList", v1.GetUser_paper_answerList)  // 获取User_paper_answer列表
	}
}
