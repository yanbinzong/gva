package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)


//@function: CreateAttachment
//@description: 创建Attachment记录
//@param: attachment model.Attachment
//@return: err error

func CreateAttachment(attachment model.Attachment) (err error) {
	err = global.GVA_DB.Create(&attachment).Error
	return err
}


//@function: DeleteAttachment
//@description: 删除Attachment记录
//@param: attachment model.Attachment
//@return: err error

func DeleteAttachment(attachment model.Attachment) (err error) {
	err = global.GVA_DB.Delete(&attachment).Error
	return err
}


//@function: DeleteAttachmentByIds
//@description: 批量删除Attachment记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAttachmentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Attachment{},"id in ?",ids.Ids).Error
	return err
}


//@function: UpdateAttachment
//@description: 更新Attachment记录
//@param: attachment *model.Attachment
//@return: err error

func UpdateAttachment(attachment model.Attachment) (err error) {
	err = global.GVA_DB.Save(&attachment).Error
	return err
}


//@function: GetAttachment
//@description: 根据id获取Attachment记录
//@param: id uint
//@return: err error, attachment model.Attachment

func GetAttachment(id uint) (err error, attachment model.Attachment) {
	err = global.GVA_DB.Where("id = ?", id).First(&attachment).Error
	return
}


//@function: GetAttachmentInfoList
//@description: 分页获取Attachment记录
//@param: info request.AttachmentSearch
//@return: err error, list interface{}, total int64

func GetAttachmentInfoList(info request.AttachmentSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Attachment{})
    var attachments []model.Attachment
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&attachments).Error
	return err, attachments, total
}