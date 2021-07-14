// 自动生成模板Attachment
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Attachment struct {
      global.GVA_MODEL
      Type  string `json:"type" form:"type" gorm:"column:type;comment:;type:varchar(191);size:191;"`
      Url  string `json:"url" form:"url" gorm:"column:url;comment:;type:tinytext;"`
      Status  int `json:"status" form:"status" gorm:"column:status;comment:;type:int;"`
      Exts  string `json:"exts" form:"exts" gorm:"column:exts;comment:;type:varchar(191);size:191;"`
      Name string `json:"name" gorm:"comment:文件名"` // 文件名
      Key  string `json:"key" gorm:"comment:编号"`   // 编号
}


func (Attachment) TableName() string {
  return "attachment"
}

