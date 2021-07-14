package v1

import (
	"fmt"
	"gin-vue-admin/api/app"
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags User_paper_answer
// @Summary 创建User_paper_answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.User_paper_answer true "创建User_paper_answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /User_paper_answer/createUser_paper_answer [post]
func CreateUser_paper_answer(c *gin.Context) {
	var User_paper_answer[] model.User_paper_answer
	_ = c.ShouldBindJSON(&User_paper_answer)
	if err := service.CreateUser_paper_answer(User_paper_answer); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags User_paper_answer
// @Summary 删除User_paper_answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.User_paper_answer true "删除User_paper_answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /User_paper_answer/deleteUser_paper_answer [delete]
func DeleteUser_paper_answer(c *gin.Context) {
	var User_paper_answer model.User_paper_answer
	_ = c.ShouldBindJSON(&User_paper_answer)
	if err := service.DeleteUser_paper_answer(User_paper_answer); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags User_paper_answer
// @Summary 批量删除User_paper_answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除User_paper_answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /User_paper_answer/deleteUser_paper_answerByIds [delete]
func DeleteUser_paper_answerByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteUser_paper_answerByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags User_paper_answer
// @Summary 更新User_paper_answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.User_paper_answer true "更新User_paper_answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /User_paper_answer/updateUser_paper_answer [put]
func UpdateUser_paper_answer(c *gin.Context) {
	var User_paper_answer model.User_paper_answer
	_ = c.ShouldBindJSON(&User_paper_answer)
	if err := service.UpdateUser_paper_answer(User_paper_answer); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags User_paper_answer
// @Summary 用id查询User_paper_answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.User_paper_answer true "用id查询User_paper_answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /user_paper_answer/findUser_paper_answer [get]
func FindUser_paper_answer(c *gin.Context) {
	var User_paper_answer model.User_paper_answer
	_ = c.ShouldBindQuery(&User_paper_answer)
	if err, reuser_paper_answer := service.GetUser_paper_answer(User_paper_answer.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuser_paper_answer": reuser_paper_answer}, c)
	}
}

// @Tags User_paper_answer
// @Summary 分页获取User_paper_answer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.User_paper_answerSearch true "分页获取User_paper_answer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /User_paper_answer/getUser_paper_answerList [get]
func GetUser_paper_answerList(c *gin.Context) {
	fmt.Println("123")
	var pageInfo request.User_paper_answerSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetUser_paper_answerInfoList(pageInfo); err != nil {
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

// @Tags User_paper_answer
// @Summary 分页获取User_paper_answer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.User_paper_answerSearch true "分页获取User_paper_answer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /User_paper_answer/Getwx_user_paper_answerList [get]
func Getwx_user_paper_answerList(c *gin.Context) {
	var requestPaper app.Request_paper
	c.ShouldBindQuery(&requestPaper)
	// service 服务层
	fmt.Println(requestPaper.Paper_id,requestPaper.Paper_id)

	if err,list_paper := service.GetQuqestAnswerList(requestPaper.Paper_id,requestPaper.User_id); err != nil{
		global.GVA_LOG.Error("暂无数据",zap.Any("err", err))
		response.FailWithMessage("暂无数据",c)
	}else{
		if(list_paper.ID==0){
			response.OkWithMessage("没有创建试卷",c)
		}else{
			response.OkWithData(gin.H{"list_paper": list_paper},c)
		}
	}



	//var pageInfo request.User_paper_answerSearch
	//_ = c.ShouldBindQuery(&pageInfo)
	//if err, list, total := service.GetUser_paper_answerInfoList(pageInfo); err != nil {
	//	global.GVA_LOG.Error("获取失败", zap.Any("err", err))
	//	response.FailWithMessage("获取失败", c)
	//} else {
	//	response.OkWithDetailed(response.PageResult{
	//		List:     list,
	//		Total:    total,
	//		Page:     pageInfo.Page,
	//		PageSize: pageInfo.PageSize,
	//	}, "获取成功", c)
	//}





}

// @Tags User_paper_answer
// @Summary 分页获取User_paper_answer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.User_paper_answerSearch true "分页获取User_paper_answer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /User_paper_answer/Getwx_user_paper_answerListAll [get]
func Getwx_user_paper_answerListAll(c *gin.Context) {


	if err, list_paper := service.GetQuqestAnswerListAll(); err != nil {
		global.GVA_LOG.Error("暂无数据", zap.Any("err", err))
		response.FailWithMessage("暂无数据", c)
	} else {
			response.OkWithData(gin.H{"list_paper": list_paper}, c)
		}
}