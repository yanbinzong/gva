// 自动生成模板Exam_paper
package model

import (
      "gin-vue-admin/global"
      "time"
)

// 如果含有time.Time 请自行import time包
type Exam_paper struct {
      global.GVA_MODEL
      Paper_title  string `json:"paper_title" form:"paper_title" gorm:"column:paper_title;comment:试卷标题;type:varchar(191);size:191;"`
      Paperfu_title  string `json:"paperfu_title" form:"paperfu_title" gorm:"column:paperfu_title;comment:试卷副标题;type:varchar(191);size:191;"`
      Paper_intro  string `json:"paper_intro" form:"paper_intro" gorm:"column:paper_intro;comment:试卷说明;type:varchar(191);size:191;"`
      Paper_status  int `json:"paper_status" form:"paper_status" gorm:"column:paper_status;comment:试卷状态;type:int;size:11;"`
      End_at *time.Time `json:"end_at" form:"end_at" gorm:"column:end_at;comment:结束时间"`
      Branch_office_id  int `json:"branch_office_id" form:"branch_office_id" gorm:"column:branch_office_id;comment:分支机构;type:int;size:11;"`
      // 关联模型：一对多
      Question []Question `gorm:"foreignKey:Exam_paper_id;references:ID"`
}

// 自定义结构体
type Request_paper_params struct {
      Exam_paper_id int `form:"exam_paper_id" json:"exam_paper_id"`
}
func (Exam_paper) TableName() string {
  return "exam_paper"
}

