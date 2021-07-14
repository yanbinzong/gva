package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)


//@function: CreateQuestion_options
//@description: 创建Question_options记录
//@param: question_options model.Question_options
//@return: err error

func CreateQuestion_options(question_options model.Question_options) (err error) {
	err = global.GVA_DB.Create(&question_options).Error
	return err
}


//@function: DeleteQuestion_options
//@description: 删除Question_options记录
//@param: question_options model.Question_options
//@return: err error

func DeleteQuestion_options(question_options model.Question_options) (err error) {
	err = global.GVA_DB.Delete(&question_options).Error
	return err
}


//@function: DeleteQuestion_optionsByIds
//@description: 批量删除Question_options记录
//@param: ids request.IdsReq
//@return: err error

func DeleteQuestion_optionsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Question_options{},"id in ?",ids.Ids).Error
	return err
}


//@function: UpdateQuestion_options
//@description: 更新Question_options记录
//@param: question_options *model.Question_options
//@return: err error

func UpdateQuestion_options(question_options model.Question_options) (err error) {
	err = global.GVA_DB.Save(&question_options).Error
	return err
}


//@function: GetQuestion_options
//@description: 根据id获取Question_options记录
//@param: id uint
//@return: err error, question_options model.Question_options

func GetQuestion_options(id uint) (err error, question_options model.Question_options) {
	err = global.GVA_DB.Where("id = ?", id).First(&question_options).Error
	return
}


//@function: GetQuestion_optionsInfoList
//@description: 分页获取Question_options记录
//@param: info request.Question_optionsSearch
//@return: err error, list interface{}, total int64

func GetQuestion_optionsInfoList(info request.Question_optionsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Question_options{})
    var question_optionss []model.Question_options
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&question_optionss).Error
	return err, question_optionss, total
}