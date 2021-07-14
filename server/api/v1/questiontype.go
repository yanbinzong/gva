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

// @Tags Question_type
// @Summary 创建Question_type
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question_type true "创建Question_type"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /question_type/createQuestion_type [post]
func CreateQuestion_type(c *gin.Context) {
	var question_type model.Question_type
	_ = c.ShouldBindJSON(&question_type)
	if err := service.CreateQuestion_type(question_type); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Question_type
// @Summary 删除Question_type
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question_type true "删除Question_type"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /question_type/deleteQuestion_type [delete]
func DeleteQuestion_type(c *gin.Context) {
	var question_type model.Question_type
	_ = c.ShouldBindJSON(&question_type)
	if err := service.DeleteQuestion_type(question_type); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Question_type
// @Summary 批量删除Question_type
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Question_type"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /question_type/deleteQuestion_typeByIds [delete]
func DeleteQuestion_typeByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteQuestion_typeByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Question_type
// @Summary 更新Question_type
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question_type true "更新Question_type"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /question_type/updateQuestion_type [put]
func UpdateQuestion_type(c *gin.Context) {
	var question_type model.Question_type
	_ = c.ShouldBindJSON(&question_type)
	if err := service.UpdateQuestion_type(question_type); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Question_type
// @Summary 用id查询Question_type
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question_type true "用id查询Question_type"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /question_type/findQuestion_type [get]
func FindQuestion_type(c *gin.Context) {
	var question_type model.Question_type
	_ = c.ShouldBindQuery(&question_type)
	if err, requestion_type := service.GetQuestion_type(question_type.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"requestion_type": requestion_type}, c)
	}
}

// @Tags Question_type
// @Summary 分页获取Question_type列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Question_typeSearch true "分页获取Question_type列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /question_type/getQuestion_typeList [get]
func GetQuestion_typeList(c *gin.Context) {
	var pageInfo request.Question_typeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetQuestion_typeInfoList(pageInfo); err != nil {
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

// 获取 question_typelist_all 全部数据
// @Tags Question_type
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Question_typeSearch true "分页获取Question_type列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /question_type/getQuestion_typeList [get]
func GetQuestion_typeList_all(c *gin.Context)  {
	if err, list := service.GetQuestion_typeInfoList_all(); err !=nil{
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	}else{
		response.OkWithDetailed(response.ListResult{
			List: list,
		},"获取成功",c)
	}
}
