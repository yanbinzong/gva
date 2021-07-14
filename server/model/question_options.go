// 自动生成模板Question_options
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Question_options struct {
      global.GVA_MODEL
      Question_id  int `json:"question_id" form:"question_id" gorm:"column:question_id;comment:问题选项表;type:int;size:11;"`
      Question_type_id  int `json:"question_type_id" form:"question_type_id" gorm:"column:question_type_id;comment:问题类型关联;type:int;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;type:varchar(191);size:191;"`
}

func (Question_options) TableName() string {
  return "question_options"
}

