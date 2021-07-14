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

// @Tags Question
// @Summary 创建Question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question true "创建Question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /question/createQuestion [post]
func CreateQuestion(c *gin.Context) {
	var question model.Question

	_ = c.ShouldBindJSON(&question)
	fmt.Println(question)
	if err := service.CreateQuestion(question); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Question
// @Summary 删除Question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question true "删除Question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /question/deleteQuestion [delete]
func DeleteQuestion(c *gin.Context) {
	var question model.Question
	_ = c.ShouldBindJSON(&question)
	if err := service.DeleteQuestion(question); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Question
// @Summary 批量删除Question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /question/deleteQuestionByIds [delete]
func DeleteQuestionByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteQuestionByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Question
// @Summary 更新Question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question true "更新Question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /question/updateQuestion [put]
func UpdateQuestion(c *gin.Context) {
	var question model.Question
	_ = c.ShouldBindJSON(&question)
	if err := service.UpdateQuestion(question); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Question
// @Summary 用id查询Question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question true "用id查询Question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /question/findQuestion [get]
func FindQuestion(c *gin.Context) {
	var question model.Question
	_ = c.ShouldBindQuery(&question)
	if err, requestion := service.GetQuestion(question.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"requestion": requestion}, c)
	}
}

// @Tags Question
// @Summary 分页获取Question列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.QuestionSearch true "分页获取Question列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /question/getQuestionList [get]
func GetQuestionList(c *gin.Context) {

	var pageInfo request.QuestionSearch

	_ = c.ShouldBindQuery(&pageInfo)
	 fmt.Println(pageInfo)
	if err, list, total := service.GetQuestionInfoList(pageInfo); err != nil {
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
