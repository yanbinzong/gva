package app

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// http://localhost:8888/app/question/?qid=1
type Request struct {
	QID int `form:"qid"`
}

func GetQuestionList(c *gin.Context) {
	global.GVA_DB.Debug().AutoMigrate(&model.Question{})
	var req Request
	c.ShouldBindQuery(&req)
	query := global.GVA_DB.Model(&model.Question{})
	if req.QID > 0 {
		query.Where("id", req.QID)
	}
	que := []model.Question{}
	// 查处题目和其关联的选项
	err := query.Preload("Question_options").Limit(10).Find(&que).Error
	if err != nil {
		global.GVA_LOG.Error("error!", zap.Any("err", err))
		response.FailWithMessage("error", c)
	} else {
		response.OkWithDetailed(que, "success", c)
	}
}
