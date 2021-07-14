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

// @Tags Question_options
// @Summary 创建Question_options
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question_options true "创建Question_options"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /question_options/createQuestion_options [post]
func CreateQuestion_options(c *gin.Context) {
	var question_options model.Question_options
	_ = c.ShouldBindJSON(&question_options)
	if err := service.CreateQuestion_options(question_options); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Question_options
// @Summary 删除Question_options
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question_options true "删除Question_options"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /question_options/deleteQuestion_options [delete]
func DeleteQuestion_options(c *gin.Context) {
	var question_options model.Question_options
	_ = c.ShouldBindJSON(&question_options)
	if err := service.DeleteQuestion_options(question_options); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Question_options
// @Summary 批量删除Question_options
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Question_options"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /question_options/deleteQuestion_optionsByIds [delete]
func DeleteQuestion_optionsByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteQuestion_optionsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Question_options
// @Summary 更新Question_options
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question_options true "更新Question_options"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /question_options/updateQuestion_options [put]
func UpdateQuestion_options(c *gin.Context) {
	var question_options model.Question_options
	_ = c.ShouldBindJSON(&question_options)
	if err := service.UpdateQuestion_options(question_options); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Question_options
// @Summary 用id查询Question_options
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question_options true "用id查询Question_options"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /question_options/findQuestion_options [get]
func FindQuestion_options(c *gin.Context) {
	var question_options model.Question_options
	_ = c.ShouldBindQuery(&question_options)
	if err, requestion_options := service.GetQuestion_options(question_options.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"requestion_options": requestion_options}, c)
	}
}

// @Tags Question_options
// @Summary 分页获取Question_options列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Question_optionsSearch true "分页获取Question_options列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /question_options/getQuestion_optionsList [get]
func GetQuestion_optionsList(c *gin.Context) {
	var pageInfo request.Question_optionsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetQuestion_optionsInfoList(pageInfo); err != nil {
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
