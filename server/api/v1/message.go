package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateMessage(c *gin.Context)  {
	var message model.Message
	_ = c.ShouldBindJSON(&message)
	fmt.Println(message)
	if err := service.CreateMessage(message); err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Any("err", err))

		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

func CreateMessageaa(c *gin.Context)  {
	var message model.Message
	message.Title = c.Request.FormValue("title")
	message.Content = c.Request.FormValue("content")
	err := c.ShouldBindJSON(&message)

	if(err != nil){
		c.JSON(200, gin.H{
			"msg":"成功了",
			"data":message,
		})
	}else {
		c.JSON(200, gin.H{
			"msg":"成功了",
			"data":message,
		})
	}
}

func FindPaper(c *gin.Context) {

	var result []model.Question_res
	global.GVA_DB.Table("question").Select("id,question_name,direction").Scan(&result)
fmt.Println(result)
	for k,v :=range result{
		fmt.Println(k,v)
	}


	//var questions []model.Question
	//global.GVA_DB.Find(&questions)
	//for k,v :=range questions{
	//	fmt.Println(k,v,"/n/r")
	//}
	//fmt.Println(questions)

	//var question = make([]*model.Question,0)
	//global.GVA_DB.Model(&question).Find(&question)
	//fmt.Printf("%v",question)
	//
	//var exam_paper model.Exam_paper
	//_ = c.ShouldBindQuery(&exam_paper)
	//
	//if err, reexam_paper := service.GetExam_paper(exam_paper.ID); err != nil {
	//	global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
	//	response.FailWithMessage("查询失败", c)
	//} else {
	//	fmt.Println(reexam_paper)
	//	c.JSON(200,gin.H{
	//		"data":reexam_paper,
	//	})
	//	//response.OkWithData(gin.H{"reexam_paper": reexam_paper}, c)
	//}
}
