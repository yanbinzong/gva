package model

import (
	"gin-vue-admin/global"
	"github.com/satori/go.uuid"
)

type SysUser struct {
	global.GVA_MODEL
	UUID        uuid.UUID    `json:"uuid" gorm:"comment:用户UUID"`                                                    // 用户UUID
	Username    string       `json:"userName" gorm:"comment:用户登录名"`                                                 // 用户登录名
	Password    string       `json:"-"  gorm:"comment:用户登录密码"`                                                      // 用户登录密码
	NickName    string       `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                     // 用户昵称
	HeaderImg   string       `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"` // 用户头像
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"` // 用户角色ID
}


//CREATE TABLE `wx_user` (
//`id`  int(11) UNSIGNED NOT NULL AUTO_INCREMENT ,
//`openid`  char(28) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '小程序用户openid' ,
//`nickname`  varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户昵称' ,
//`avatarurl`  varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户头像' ,
//`gender`  tinyint(1) NULL DEFAULT 0 COMMENT '性别   0 男  1  女  2 人妖' ,
//`country`  varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '所在国家' ,
//`province`  varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '省份' ,
//`city`  varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '城市' ,
//`language`  varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
//`ctime`  int(11) NULL DEFAULT NULL COMMENT '创建时间' ,
//`mobile`  varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '手机类型' ,
//`telnum`  char(13) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '手机号码' ,
//PRIMARY KEY (`id`)
