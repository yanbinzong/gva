// 自动生成模板Question_type
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Question_type struct {
      global.GVA_MODEL
      Qtype_name  string `json:"qtype_name" form:"qtype_name" gorm:"column:qtype_name;comment:;type:varchar(191);size:191;"`
      Qtype_content  string `json:"qtype_content" form:"qtype_content" gorm:"column:qtype_content;comment:;type:varchar(191);size:191;"`
}


func (Question_type) TableName() string {
  return "question_type"
}

