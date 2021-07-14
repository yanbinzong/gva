package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)


//@function: CreateWx_user
//@description: 创建Wx_user记录
//@param: wx_user model.Wx_user
//@return: err error

func CreateWx_user(wx_user model.Wx_user) (err error) {
	err = global.GVA_DB.Create(&wx_user).Error
	return err
}


//@function: DeleteWx_user
//@description: 删除Wx_user记录
//@param: wx_user model.Wx_user
//@return: err error

func DeleteWx_user(wx_user model.Wx_user) (err error) {
	err = global.GVA_DB.Delete(&wx_user).Error
	return err
}


//@function: DeleteWx_userByIds
//@description: 批量删除Wx_user记录
//@param: ids request.IdsReq
//@return: err error

func DeleteWx_userByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Wx_user{},"id in ?",ids.Ids).Error
	return err
}


//@function: UpdateWx_user
//@description: 更新Wx_user记录
//@param: wx_user *model.Wx_user
//@return: err error

func UpdateWx_user(wx_user model.Wx_user) (err error) {
	err = global.GVA_DB.Save(&wx_user).Error
	return err
}


//@function: GetWx_user
//@description: 根据id获取Wx_user记录
//@param: id uint
//@return: err error, wx_user model.Wx_user

func GetWx_user(id uint) (err error, wx_user model.Wx_user) {
	err = global.GVA_DB.Where("id = ?", id).First(&wx_user).Error
	return
}


//@function: GetWx_userInfoList
//@description: 分页获取Wx_user记录
//@param: info request.Wx_userSearch
//@return: err error, list interface{}, total int64

func GetWx_userInfoList(info request.Wx_userSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Wx_user{})
    var wx_users []model.Wx_user
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Openid != "" {
        db = db.Where("`openid` = ?",info.Openid)
    }
    if info.Nickname != "" {
        db = db.Where("`nickname` LIKE ?","%"+ info.Nickname+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&wx_users).Error
	return err, wx_users, total
}

//@function: GetWx_user
//@description: 根据id获取Wx_user记录
//@param: id uint
//@return: err error, wx_user model.Wx_user

func GetWx_user_openid(openid string) (err error, wx_user model.Wx_user) {
	err = global.GVA_DB.Where("openid = ?", openid).First(&wx_user).Error
	return
}
