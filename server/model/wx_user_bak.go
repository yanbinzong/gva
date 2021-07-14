// 自动生成模板Wx_user
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Wx_users struct {
      global.GVA_MODEL
      Openid  string `json:"openid" form:"openid" gorm:"column:openid;comment:;type:char(30);size:30;"`
      Nickname  string `json:"nickname" form:"nickname" gorm:"column:nickname;comment:;type:varchar(100);size:100;"`
      Avatarurl  string `json:"avatarurl" form:"avatarurl" gorm:"column:avatarurl;comment:;type:varchar(1000);size:1000;"`
      Gender  string `json:"gender" form:"gender" gorm:"column:gender;comment:;type:varchar;"`
      Country  string `json:"country" form:"country" gorm:"column:country;comment:;type:varchar(100);size:100;"`
      Province  string `json:"province" form:"province" gorm:"column:province;comment:;type:varchar(100);size:100;"`
      City  string `json:"city" form:"city" gorm:"column:city;comment:;type:varchar(100);size:100;"`
      Language  string `json:"language" form:"language" gorm:"column:language;comment:;type:varchar(100);size:100;"`
      Mobile  string `json:"mobile" form:"mobile" gorm:"column:mobile;comment:;type:varchar(50);size:50;"`
}


func (Wx_users) TableName() string {
  return "wx_users"
}

