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

// @Tags Attachment
// @Summary 创建Attachment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Attachment true "创建Attachment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /attachment/createAttachment [post]
func CreateAttachment(c *gin.Context) {
	var attachment model.Attachment
	_ = c.ShouldBindJSON(&attachment)
	if err := service.CreateAttachment(attachment); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Attachment
// @Summary 删除Attachment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Attachment true "删除Attachment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /attachment/deleteAttachment [delete]
func DeleteAttachment(c *gin.Context) {
	var attachment model.Attachment
	_ = c.ShouldBindJSON(&attachment)
	if err := service.DeleteAttachment(attachment); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Attachment
// @Summary 批量删除Attachment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Attachment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /attachment/deleteAttachmentByIds [delete]
func DeleteAttachmentByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteAttachmentByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Attachment
// @Summary 更新Attachment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Attachment true "更新Attachment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /attachment/updateAttachment [put]
func UpdateAttachment(c *gin.Context) {
	var attachment model.Attachment
	_ = c.ShouldBindJSON(&attachment)
	if err := service.UpdateAttachment(attachment); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Attachment
// @Summary 用id查询Attachment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Attachment true "用id查询Attachment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /attachment/findAttachment [get]
func FindAttachment(c *gin.Context) {
	var attachment model.Attachment
	_ = c.ShouldBindQuery(&attachment)
	if err, reattachment := service.GetAttachment(attachment.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reattachment": reattachment}, c)
	}
}

// @Tags Attachment
// @Summary 分页获取Attachment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AttachmentSearch true "分页获取Attachment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /attachment/getAttachmentList [get]
func GetAttachmentList(c *gin.Context) {
	var pageInfo request.AttachmentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetAttachmentInfoList(pageInfo); err != nil {
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
