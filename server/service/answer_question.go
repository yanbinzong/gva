package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)


//@function: CreateAnswer_question
//@description: 创建Answer_question记录
//@param: answer_question model.Answer_question
//@return: err error

func CreateAnswer_question(answer_question model.Answer_question) (err error) {
	err = global.GVA_DB.Create(&answer_question).Error
	return err
}


//@function: DeleteAnswer_question
//@description: 删除Answer_question记录
//@param: answer_question model.Answer_question
//@return: err error

func DeleteAnswer_question(answer_question model.Answer_question) (err error) {
	err = global.GVA_DB.Delete(&answer_question).Error
	return err
}


//@function: DeleteAnswer_questionByIds
//@description: 批量删除Answer_question记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAnswer_questionByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Answer_question{},"id in ?",ids.Ids).Error
	return err
}


//@function: UpdateAnswer_question
//@description: 更新Answer_question记录
//@param: answer_question *model.Answer_question
//@return: err error

func UpdateAnswer_question(answer_question model.Answer_question) (err error) {
	err = global.GVA_DB.Save(&answer_question).Error
	return err
}


//@function: GetAnswer_question
//@description: 根据id获取Answer_question记录
//@param: id uint
//@return: err error, answer_question model.Answer_question

func GetAnswer_question(id uint) (err error, answer_question model.Answer_question) {
	err = global.GVA_DB.Where("id = ?", id).First(&answer_question).Error
	return
}


//@function: GetAnswer_questionInfoList
//@description: 分页获取Answer_question记录
//@param: info request.Answer_questionSearch
//@return: err error, list interface{}, total int64

func GetAnswer_questionInfoList(info request.Answer_questionSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Answer_question{})
    var answer_questions []model.Answer_question
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.User_id != 0 {
        db = db.Where("`user_id` = ?",info.User_id)
    }
    if info.Paper_id != 0 {
        db = db.Where("`paper_id` = ?",info.Paper_id)
    }
    if info.Question_id != 0 {
        db = db.Where("`question_id` = ?",info.Question_id)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&answer_questions).Error
	return err, answer_questions, total
}