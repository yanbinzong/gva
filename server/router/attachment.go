package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAttachmentRouter(Router *gin.RouterGroup) {
	AttachmentRouter := Router.Group("attachment").Use(middleware.OperationRecord())
	{
		AttachmentRouter.POST("createAttachment", v1.CreateAttachment)   // 新建Attachment
		AttachmentRouter.DELETE("deleteAttachment", v1.DeleteAttachment) // 删除Attachment
		AttachmentRouter.DELETE("deleteAttachmentByIds", v1.DeleteAttachmentByIds) // 批量删除Attachment
		AttachmentRouter.PUT("updateAttachment", v1.UpdateAttachment)    // 更新Attachment
		AttachmentRouter.GET("findAttachment", v1.FindAttachment)        // 根据ID获取Attachment
		AttachmentRouter.GET("getAttachmentList", v1.GetAttachmentList)  // 获取Attachment列表
	}
}
