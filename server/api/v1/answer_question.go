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

// @Tags Answer_question
// @Summary 创建Answer_question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Answer_question true "创建Answer_question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /answer_question/createAnswer_question [post]
func CreateAnswer_question(c *gin.Context) {
	var answer_question model.Answer_question
	_ = c.ShouldBindJSON(&answer_question)
	if err := service.CreateAnswer_question(answer_question); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Answer_question
// @Summary 删除Answer_question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Answer_question true "删除Answer_question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /answer_question/deleteAnswer_question [delete]
func DeleteAnswer_question(c *gin.Context) {
	var answer_question model.Answer_question
	_ = c.ShouldBindJSON(&answer_question)
	if err := service.DeleteAnswer_question(answer_question); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Answer_question
// @Summary 批量删除Answer_question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Answer_question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /answer_question/deleteAnswer_questionByIds [delete]
func DeleteAnswer_questionByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteAnswer_questionByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Answer_question
// @Summary 更新Answer_question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Answer_question true "更新Answer_question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /answer_question/updateAnswer_question [put]
func UpdateAnswer_question(c *gin.Context) {
	var answer_question model.Answer_question
	_ = c.ShouldBindJSON(&answer_question)
	if err := service.UpdateAnswer_question(answer_question); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Answer_question
// @Summary 用id查询Answer_question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Answer_question true "用id查询Answer_question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /answer_question/findAnswer_question [get]
func FindAnswer_question(c *gin.Context) {
	var answer_question model.Answer_question
	_ = c.ShouldBindQuery(&answer_question)
	if err, reanswer_question := service.GetAnswer_question(answer_question.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reanswer_question": reanswer_question}, c)
	}
}

// @Tags Answer_question
// @Summary 分页获取Answer_question列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Answer_questionSearch true "分页获取Answer_question列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /answer_question/getAnswer_questionList [get]
func GetAnswer_questionList(c *gin.Context) {
	var pageInfo request.Answer_questionSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetAnswer_questionInfoList(pageInfo); err != nil {
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
