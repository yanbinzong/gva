package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitBranch_officeRouter(Router *gin.RouterGroup) {
	Branch_officeRouter := Router.Group("branch_office").Use(middleware.OperationRecord())
	{
		Branch_officeRouter.POST("createBranch_office", v1.CreateBranch_office)   // 新建Branch_office
		Branch_officeRouter.DELETE("deleteBranch_office", v1.DeleteBranch_office) // 删除Branch_office
		Branch_officeRouter.DELETE("deleteBranch_officeByIds", v1.DeleteBranch_officeByIds) // 批量删除Branch_office
		Branch_officeRouter.PUT("updateBranch_office", v1.UpdateBranch_office)    // 更新Branch_office
		Branch_officeRouter.GET("findBranch_office", v1.FindBranch_office)        // 根据ID获取Branch_office
		Branch_officeRouter.GET("getBranch_officeList", v1.GetBranch_officeList)  // 获取Branch_office列表
	}
}
