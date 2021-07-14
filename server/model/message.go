// 自动生成模板Question
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Message struct {
	global.GVA_MODEL
	//Title  string `json:"title" form:"title" gorm:"column:title;comment:消息标题;type:varchar(191);size:191;"`
	//Content  string `json:"content" form:"content" gorm:"column:content;comment:消息内容;type:varchar(191);size:191;"`
	//Answer_state  int `json:"answer" form:"answer" gorm:"column:answer;comment:是否必答"`
	//Topic_id  string `json:"topic_id" form:"topic_id" gorm:"column:topic_id;comment:题型编号;type:varchar(191);size:191;"`
	//Author  string `json:"author" form:"author" gorm:"column:author;comment:发题人;type:varchar(191);size:191;"`
	Title string `form:"title" json:"title"`
	Content string `form:"content" json:"content"`

}


func (Message) TableName() string {
	return "message"
}
