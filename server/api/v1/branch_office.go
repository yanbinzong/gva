package v1

import (
	"fmt"
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags Branch_office
// @Summary 创建Branch_office
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Branch_office true "创建Branch_office"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /branch_office/createBranch_office [post]
func CreateBranch_office(c *gin.Context) {
	var branch_office model.Branch_office
	_ = c.ShouldBindJSON(&branch_office)
	if err := service.CreateBranch_office(branch_office); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Branch_office
// @Summary 删除Branch_office
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Branch_office true "删除Branch_office"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /branch_office/deleteBranch_office [delete]
func DeleteBranch_office(c *gin.Context) {
	var branch_office model.Branch_office
	_ = c.ShouldBindJSON(&branch_office)
	if err := service.DeleteBranch_office(branch_office); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Branch_office
// @Summary 批量删除Branch_office
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Branch_office"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /branch_office/deleteBranch_officeByIds [delete]
func DeleteBranch_officeByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteBranch_officeByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Branch_office
// @Summary 更新Branch_office
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Branch_office true "更新Branch_office"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /branch_office/updateBranch_office [put]
func UpdateBranch_office(c *gin.Context) {
	var branch_office model.Branch_office
	_ = c.ShouldBindJSON(&branch_office)
	if err := service.UpdateBranch_office(branch_office); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Branch_office
// @Summary 用id查询Branch_office
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Branch_office true "用id查询Branch_office"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /branch_office/findBranch_office [get]
func FindBranch_office(c *gin.Context) {
	var branch_office model.Branch_office
	_ = c.ShouldBindQuery(&branch_office)
	if err, rebranch_office := service.GetBranch_office(branch_office.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebranch_office": rebranch_office}, c)
	}
}

// @Tags Branch_office
// @Summary 分页获取Branch_office列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Branch_officeSearch true "分页获取Branch_office列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /branch_office/getBranch_officeList [get]
func GetBranch_officeList(c *gin.Context) {
	var pageInfo request.Branch_officeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetBranch_officeInfoList(pageInfo); err != nil {  //_list 先把原来的列表隐藏掉
	   global.GVA_LOG.Error("获取失败", zap.Any("err", err))
       response.FailWithMessage("获取失败", c)
    } else {
       response.OkWithDetailed(response.PageResult{
           List:     list,
           Total:    total,
           Page:     pageInfo.Page,
           PageSize: pageInfo.PageSize,
       }, "获取成功", c)
    }
}

/*
treeList := getMenu(0)
递归获取树形菜单
*/
func getMenu(pid uint) [] *model.TreeList {
	var Branch_office []model.Branch_office
	global.GVA_DB.Where("pid = ?", pid).Find(&Branch_office) // 拿到所有父菜单(pid==0)
	treeList := []*model.TreeList{}
	for _, v := range Branch_office { // 循环所有父菜单
		fmt.Println(v.ID)
		child := getMenu(v.ID) // 拿到每个父菜单的子菜单
		node := &model.TreeList{ // 拼装父菜单数据
			Name:  v.Name,
			Stataus:  v.Stataus,
			Remarks:   v.Remarks,
			Pid: v.Pid,
		}
		node.Children  = child
		treeList = append(treeList, node)
	}
	return treeList
}

