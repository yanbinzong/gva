// 自动生成模板Wx_user
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Wx_user struct {
      global.GVA_MODEL
      Openid  string `json:"openid" form:"openid" gorm:"column:openid;comment:;type:varchar(30);size:30;"`
      Nickname  string `json:"nickname" form:"nickname" gorm:"column:nickname;comment:;type:varchar(100);size:100;"`
      Avatarurl  string `json:"avatarurl" form:"avatarurl" gorm:"column:avatarurl;comment:用户头像;type:varchar(120);size:120;"`
      Gender  string `json:"gender" form:"gender" gorm:"column:gender;comment:性别;type:varchar(10);size:10;"`
      Country  string `json:"country" form:"country" gorm:"column:country;comment:所在国家;type:varchar(100);size:100;"`
      Province  string `json:"province" form:"province" gorm:"column:province;comment:省份;type:varchar(100);size:100;"`
      City  string `json:"city" form:"city" gorm:"column:city;comment:城市;type:varchar(100);size:100;"`
      Language  string `json:"language" form:"language" gorm:"column:language;comment:语言;type:varchar(100);size:100;"`
      Mobile  string `json:"mobile" form:"mobile" gorm:"column:mobile;comment:手机号码;type:varchar(50);size:50;"`
      SessionKey  string `json:"sessionKey" form:"sessionKey" gorm:"column:session_key;comment:sessionKey;type:varchar(40);size:40;"`
      MySession  string `json:"mySession" form:"mySession" gorm:"column:my_session;comment:mySession;type:varchar(40);size:40;"`
}

func (Wx_user) TableName() string {
  return "wx_user"
}

