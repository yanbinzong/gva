package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags Wx_user
// @Summary 创建Wx_user
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Wx_user true "创建Wx_user"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wx_user/createWx_user [post]
func CreateWx_user(c *gin.Context) {
	var wx_user model.Wx_user
	_ = c.ShouldBindJSON(&wx_user)
	if err := service.CreateWx_user(wx_user); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Wx_user
// @Summary 删除Wx_user
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Wx_user true "删除Wx_user"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wx_user/deleteWx_user [delete]
func DeleteWx_user(c *gin.Context) {
	var wx_user model.Wx_user
	_ = c.ShouldBindJSON(&wx_user)
	if err := service.DeleteWx_user(wx_user); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Wx_user
// @Summary 批量删除Wx_user
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Wx_user"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wx_user/deleteWx_userByIds [delete]
func DeleteWx_userByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteWx_userByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Wx_user
// @Summary 更新Wx_user
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Wx_user true "更新Wx_user"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wx_user/updateWx_user [put]
func UpdateWx_user(c *gin.Context) {
	var wx_user model.Wx_user
	_ = c.ShouldBindJSON(&wx_user)
	if err := service.UpdateWx_user(wx_user); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Wx_user
// @Summary 用id查询Wx_user
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Wx_user true "用id查询Wx_user"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wx_user/findWx_user [get]
func FindWx_user(c *gin.Context) {
	var wx_user model.Wx_user
	_ = c.ShouldBindQuery(&wx_user)
	if err, rewx_user := service.GetWx_user(wx_user.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewx_user": rewx_user}, c)
	}
}

// @Tags Wx_user
// @Summary 分页获取Wx_user列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Wx_userSearch true "分页获取Wx_user列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wx_user/getWx_userList [get]
func GetWx_userList(c *gin.Context) {
	var pageInfo request.Wx_userSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetWx_userInfoList(pageInfo); err != nil {
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

// 2021 7 8
// @Tags Wx_user
// @Summary 用id查询Wx_user
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Wx_user true "用openid查询Wx_user"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wx_user/findWx_user_openid [get]
func FindWx_user_openid(c *gin.Context) {
	var wx_user model.Wx_user
	_ = c.ShouldBindQuery(&wx_user)
	if err, rewx_user := service.GetWx_user_openid(wx_user.Openid); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewx_user": rewx_user}, c)
	}
}
