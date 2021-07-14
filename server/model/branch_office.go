// 自动生成模板Branch_office
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Branch_office struct {
      global.GVA_MODEL
      Pid  int `json:"pid" form:"pid" gorm:"column:pid;comment:父ID;type:int;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名称;type:varchar(191);size:191;"`
      Stataus  int `json:"status" form:"status" gorm:"column:status;comment:状态;type:int;"`
      Remarks  string `json:"remarks" form:"remarks" gorm:"column:remarks;comment:备注;type:text;"`
}


// 定义一个序列化数据的结构体
type TreeList struct {
      global.GVA_MODEL
      Pid         int        `json:"pid"`
      Name        string    `json:"name"`
      Stataus     int       `json:"stataus"`
      Remarks     string    `json:"remarks"`
      Children  []*TreeList `json:"children"`
}

// 定义一个response 结构体 用于构造返回数据
type DataRes struct {
      List   []*TreeList   `json:"list"`
}
func (Branch_office) TableName() string {
  return "branch_office"
}

