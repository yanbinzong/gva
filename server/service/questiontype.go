package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)


//@function: CreateQuestion_type
//@description: 创建Question_type记录
//@param: question_type model.Question_type
//@return: err error

func CreateQuestion_type(question_type model.Question_type) (err error) {
	err = global.GVA_DB.Create(&question_type).Error
	return err
}


//@function: DeleteQuestion_type
//@description: 删除Question_type记录
//@param: question_type model.Question_type
//@return: err error

func DeleteQuestion_type(question_type model.Question_type) (err error) {
	err = global.GVA_DB.Delete(&question_type).Error
	return err
}


//@function: DeleteQuestion_typeByIds
//@description: 批量删除Question_type记录
//@param: ids request.IdsReq
//@return: err error

func DeleteQuestion_typeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Question_type{},"id in ?",ids.Ids).Error
	return err
}


//@function: UpdateQuestion_type
//@description: 更新Question_type记录
//@param: question_type *model.Question_type
//@return: err error

func UpdateQuestion_type(question_type model.Question_type) (err error) {
	err = global.GVA_DB.Save(&question_type).Error
	return err
}


//@function: GetQuestion_type
//@description: 根据id获取Question_type记录
//@param: id uint
//@return: err error, question_type model.Question_type

func GetQuestion_type(id uint) (err error, question_type model.Question_type) {
	err = global.GVA_DB.Where("id = ?", id).First(&question_type).Error
	return
}


//@function: GetQuestion_typeInfoList
//@description: 分页获取Question_type记录
//@param: info request.Question_typeSearch
//@return: err error, list interface{}, total int64

func GetQuestion_typeInfoList(info request.Question_typeSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Question_type{})
    var question_types []model.Question_type
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&question_types).Error
	return err, question_types, total
}

// GetQuestion_typeInfoList_all
func GetQuestion_typeInfoList_all()(err error,list interface{})  {
	var question_type []model.Question_type
	err = global.GVA_DB.Find(&question_type).Error
	return err,question_type
}
