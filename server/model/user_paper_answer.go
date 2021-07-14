// 自动生成模板User_paper_answer
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type User_paper_answer struct {
      global.GVA_MODEL
      User_id  int `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户ID;type:int;"`
      Paper_id  int `json:"paper_id" form:"paper_id" gorm:"column:paper_id;comment:试卷ID;type:int;"`
      Question_id  int `json:"question_id" form:"question_id" gorm:"column:question_id;comment:问题ID"`
      Answer_content  string `json:"answer_content" form:"answer_content" gorm:"column:answer_content;comment:回答内容;type:tinytext;"`
      Score  float64 `json:"score" form:"score" gorm:"column:score;comment:分数;type:decimal;"`
}


func (User_paper_answer) TableName() string {
  return "User_paper_answer"
}

