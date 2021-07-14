package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)


//@function: CreateQuestion
//@description: 创建Question记录
//@param: question model.Question
//@return: err error

func CreateQuestion(question model.Question) (err error) {
	err = global.GVA_DB.Create(&question).Error
	return err
}


//@function: DeleteQuestion
//@description: 删除Question记录
//@param: question model.Question
//@return: err error

func DeleteQuestion(question model.Question) (err error) {
	err = global.GVA_DB.Delete(&question).Error
	return err
}


//@function: DeleteQuestionByIds
//@description: 批量删除Question记录
//@param: ids request.IdsReq
//@return: err error

func DeleteQuestionByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Question{},"id in ?",ids.Ids).Error
	return err
}


//@function: UpdateQuestion
//@description: 更新Question记录
//@param: question *model.Question
//@return: err error

func UpdateQuestion(question model.Question) (err error) {
	err = global.GVA_DB.Save(&question).Error
	return err
}


//@function: GetQuestion
//@description: 根据id获取Question记录
//@param: id uint
//@return: err error, question model.Question

func GetQuestion(id uint) (err error, question model.Question) {
	err = global.GVA_DB.Where("id = ?", id).First(&question).Error
	return
}


//@function: GetQuestionInfoList
//@description: 分页获取Question记录
//@param: info request.QuestionSearch
//@return: err error, list interface{}, total int64

func GetQuestionInfoList(info request.QuestionSearch) (err error, list interface{}, total int64) {
	fmt.Println(info.Exam_paper_id)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Debug().Model(&model.Question{})
    var questions []model.Question
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Question_name != "" {
        db = db.Where("`question_name` LIKE ?","%"+ info.Question_name+"%")
    }
    if info.Answer_state != 0 {
        db = db.Where("`answer` = ?",info.Answer_state)
    }
    if info.Topic_id != "" {
        db = db.Where("`topic_id` = ?",info.Topic_id)
    }
    if info.Author != "" {
        db = db.Where("`author` LIKE ?","%"+ info.Author+"%")
    }
    if info.Exam_paper_id != 0 {
		db = db.Where("`exam_paper_id` = ?",info.Exam_paper_id)
	}
	//db = db.Where("`exam_paper_id` = ?",1)
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("sort ASC").Order("ID DESC").Find(&questions).Error
	return err, questions, total
}

//@function: exam_paper_id
//@description: 根据exam_paper_id获取Question记录
//@param: id uint
//@return: err error, question model.Question
func GetPaperQuestion(exam_paper_id int)(err error,question model.Question)  {
	fmt.Println(exam_paper_id)
	err = global.GVA_DB.Where("exam_paper_id = ?",exam_paper_id).Find(question).Error
	return
}


