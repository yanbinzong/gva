package router

import (
	"gin-vue-admin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitWx_userRouter(Router *gin.RouterGroup) {
	//Wx_userRouter := Router.Group("wx_user").Use(middleware.OperationRecord())
	Wx_userRouter := Router.Group("wx_user")
	{
		Wx_userRouter.POST("createWx_user", v1.CreateWx_user)   // 新建Wx_user
		Wx_userRouter.DELETE("deleteWx_user", v1.DeleteWx_user) // 删除Wx_user
		Wx_userRouter.DELETE("deleteWx_userByIds", v1.DeleteWx_userByIds) // 批量删除Wx_user
		Wx_userRouter.PUT("updateWx_user", v1.UpdateWx_user)    // 更新Wx_user
		Wx_userRouter.GET("findWx_user", v1.FindWx_user)        // 根据ID获取Wx_user
		Wx_userRouter.GET("getWx_userList", v1.GetWx_userList)  // 获取Wx_user列表
		Wx_userRouter.GET("findWx_user_openid",v1.FindWx_user_openid)        // 根据openid获取Wx_user
	}


	Wx_user_paper_answerRouter := Router.Group("wx_user_answer")
	{
		Wx_user_paper_answerRouter.POST("createUser_paper_answer", v1.CreateUser_paper_answer)   // 新建User_paper_answer
		Wx_user_paper_answerRouter.DELETE("deleteUser_paper_answer", v1.DeleteUser_paper_answer) // 删除User_paper_answer
		Wx_user_paper_answerRouter.DELETE("deleteUser_paper_answerByIds", v1.DeleteUser_paper_answerByIds) // 批量删除User_paper_answer
		Wx_user_paper_answerRouter.PUT("updateUser_paper_answer", v1.UpdateUser_paper_answer)    // 更新User_paper_answer
		Wx_user_paper_answerRouter.GET("findUser_paper_answer", v1.FindUser_paper_answer)        // 根据ID获取User_paper_answer
		Wx_user_paper_answerRouter.GET("getwxuser_paper_answerList", v1.Getwx_user_paper_answerList)  // 获取User_paper_answer列表
		Wx_user_paper_answerRouter.GET("getwxuser_paper_answerListAll", v1.Getwx_user_paper_answerListAll)
	}

}
