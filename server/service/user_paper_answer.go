package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)


//@function: CreateUser_paper_answer
//@description: 创建User_paper_answer记录
//@param: User_paper_answer model.User_paper_answer
//@return: err error

func CreateUser_paper_answer(user_paper_answer[] model.User_paper_answer) (err error) {
	global.GVA_DB.AutoMigrate(&user_paper_answer)
	err = global.GVA_DB.Create(&user_paper_answer).Error
	return err
}


//@function: DeleteUser_paper_answer
//@description: 删除User_paper_answer记录
//@param: User_paper_answer model.User_paper_answer
//@return: err error

func DeleteUser_paper_answer(user_paper_answer model.User_paper_answer) (err error) {
	err = global.GVA_DB.Delete(&user_paper_answer).Error
	return err
}


//@function: DeleteUser_paper_answerByIds
//@description: 批量删除User_paper_answer记录
//@param: ids request.IdsReq
//@return: err error

func DeleteUser_paper_answerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.User_paper_answer{},"id in ?",ids.Ids).Error
	return err
}


//@function: UpdateUser_paper_answer
//@description: 更新User_paper_answer记录
//@param: User_paper_answer *model.User_paper_answer
//@return: err error

func UpdateUser_paper_answer(user_paper_answer model.User_paper_answer) (err error) {
	err = global.GVA_DB.Save(&user_paper_answer).Error
	return err
}


//@function: GetUser_paper_answer
//@description: 根据id获取User_paper_answer记录
//@param: id uint
//@return: err error, User_paper_answer model.User_paper_answer

func GetUser_paper_answer(id uint) (err error, user_paper_answer model.User_paper_answer) {
	err = global.GVA_DB.Where("id = ?", id).First(&user_paper_answer).Error
	return
}


//@function: GetUser_paper_answerInfoList
//@description: 分页获取User_paper_answer记录
//@param: info request.User_paper_answerSearch
//@return: err error, list interface{}, total int64

func GetUser_paper_answerInfoList(info request.User_paper_answerSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.User_paper_answer{})
    var user_paper_answer []model.User_paper_answer
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&user_paper_answer).Error
	return err, user_paper_answer, total
}