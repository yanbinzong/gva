package router

import (
	"gin-vue-admin/api/app"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitAppRouter(r *gin.RouterGroup) {
	q := r.Group("question")
	{
		q.GET("/", app.GetQuestionList) // 合并文件
	}

	PaperRouter :=r.Group("paper")
	{
		PaperRouter.GET("/",app.GetPaperList) //获取试卷ID相关信列表
	}

	// 图片上传 扩展 extend
	upload := r.Group("extend")
	{
		upload.POST("/upload",app.UploadFile)                                 // 上传文件
	}

	Web_user_paper_answerRouter := r.Group("answer_question")
	{
		Web_user_paper_answerRouter.POST("createAnswer_question_front", app.CreateUser_paper_answer)
		Web_user_paper_answerRouter.GET("")
	}

	// 微信登录
	WeChat := r.Group("wechat")
	// 基于cookie创建session的存储引擎，传递一个参数，用来做加密时的密钥
	store := cookie.NewStore([]byte("secret11111"))
	//session中间件生效，参数mysession，是浏览器端cookie的名字
	WeChat.Use(sessions.Sessions("mysession", store))
	{
		WeChat.GET("applet_login",app.AppletWeChatLogin)
	}

}

