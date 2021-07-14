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

// @Tags Exam_paper
// @Summary 创建Exam_paper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Exam_paper true "创建Exam_paper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exam_paper/createExam_paper [post]
func CreateExam_paper(c *gin.Context) {
	var exam_paper model.Exam_paper
	_ = c.ShouldBindJSON(&exam_paper)
	if err := service.CreateExam_paper(exam_paper); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Exam_paper
// @Summary 删除Exam_paper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Exam_paper true "删除Exam_paper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exam_paper/deleteExam_paper [delete]
func DeleteExam_paper(c *gin.Context) {
	var exam_paper model.Exam_paper
	_ = c.ShouldBindJSON(&exam_paper)
	if err := service.DeleteExam_paper(exam_paper); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Exam_paper
// @Summary 批量删除Exam_paper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Exam_paper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /exam_paper/deleteExam_paperByIds [delete]
func DeleteExam_paperByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteExam_paperByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Exam_paper
// @Summary 更新Exam_paper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Exam_paper true "更新Exam_paper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exam_paper/updateExam_paper [put]
func UpdateExam_paper(c *gin.Context) {
	var exam_paper model.Exam_paper
	_ = c.ShouldBindJSON(&exam_paper)
	fmt.Println(exam_paper)
	if err := service.UpdateExam_paper(exam_paper); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Exam_paper
// @Summary 用id查询Exam_paper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Exam_paper true "用id查询Exam_paper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exam_paper/findExam_paper [get]
func FindExam_paper(c *gin.Context) {
	var exam_paper model.Exam_paper
	_ = c.ShouldBindQuery(&exam_paper)
	if err, reexam_paper := service.GetExam_paper(exam_paper.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reexam_paper": reexam_paper}, c)
	}
}

// @Tags Exam_paper
// @Summary 分页获取Exam_paper列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Exam_paperSearch true "分页获取Exam_paper列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exam_paper/getExam_paperList [get]
func GetExam_paperList(c *gin.Context) {
	var pageInfo request.Exam_paperSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetExam_paperInfoList(pageInfo); err != nil {
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



// @Tags Exam_paper
// @Summary 获取试卷->全部问题
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Request_paper_params true "试卷全部列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exam_paper/FindExam_paperQuestion [get]
func FindExam_paperQuestion(c *gin.Context)  {
	 var exam_paper_id model.Request_paper_params
	//var requestPaper app.Request_paper
	c.ShouldBindQuery(&exam_paper_id) // 获取试卷ID
	fmt.Println(exam_paper_id.Exam_paper_id)
	if err, list_paper :=service.GetQuestList(exam_paper_id.Exam_paper_id); err != nil{
		response.OkWithMessage("没有创建试卷",c)
	}else{
		response.OkWithData(gin.H{"list": list_paper},c)
	}
}

//
//func FindExam_paperQuestion1(c *gin.Context)  {
//	var request_paper_params model.Request_paper_params
//	c.ShouldBindQuery(&request_paper_params) // 获取试卷ID
//	fmt.Println(request_paper_params.Exam_paper_id)
//	if err, list_paper :=service.GetPaperQuestion(request_paper_params.Exam_paper_id); err != nil{
//		response.OkWithMessage("没有创建试卷",c)
//	}else{
//		response.OkWithData(gin.H{"list_paper": list_paper},c)
//	}
//}
