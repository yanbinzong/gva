package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitExam_paperRouter(Router *gin.RouterGroup) {
	Exam_paperRouter := Router.Group("exam_paper").Use(middleware.OperationRecord())
	{
		Exam_paperRouter.POST("createExam_paper", v1.CreateExam_paper)   // 新建Exam_paper
		Exam_paperRouter.DELETE("deleteExam_paper", v1.DeleteExam_paper) // 删除Exam_paper
		Exam_paperRouter.DELETE("deleteExam_paperByIds", v1.DeleteExam_paperByIds) // 批量删除Exam_paper
		Exam_paperRouter.PUT("updateExam_paper", v1.UpdateExam_paper)    // 更新Exam_paper
		Exam_paperRouter.GET("findExam_paper", v1.FindExam_paper)        // 根据ID获取Exam_paper
		Exam_paperRouter.GET("getExam_paperList", v1.GetExam_paperList)  // 获取Exam_paper列表
		Exam_paperRouter.GET("findExam_paperQuestion", v1.FindExam_paperQuestion)    // 根据ID获取Exam_paper

	}
}
