package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)


//@function: CreateBranch_office
//@description: 创建Branch_office记录
//@param: branch_office model.Branch_office
//@return: err error

func CreateBranch_office(branch_office model.Branch_office) (err error) {

	global.GVA_DB.AutoMigrate(&branch_office)  //创建数据表

	err = global.GVA_DB.Create(&branch_office).Error
	return err
}


//@function: DeleteBranch_office
//@description: 删除Branch_office记录
//@param: branch_office model.Branch_office
//@return: err error

func DeleteBranch_office(branch_office model.Branch_office) (err error) {
	err = global.GVA_DB.Delete(&branch_office).Error
	return err
}


//@function: DeleteBranch_officeByIds
//@description: 批量删除Branch_office记录
//@param: ids request.IdsReq
//@return: err error

func DeleteBranch_officeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Branch_office{},"id in ?",ids.Ids).Error
	return err
}


//@function: UpdateBranch_office
//@description: 更新Branch_office记录
//@param: branch_office *model.Branch_office
//@return: err error

func UpdateBranch_office(branch_office model.Branch_office) (err error) {
	err = global.GVA_DB.Save(&branch_office).Error
	return err
}


//@function: GetBranch_office
//@description: 根据id获取Branch_office记录
//@param: id uint
//@return: err error, branch_office model.Branch_office

func GetBranch_office(id uint) (err error, branch_office model.Branch_office) {
	err = global.GVA_DB.Where("id = ?", id).First(&branch_office).Error
	return
}



//@function: GetBranch_officeInfoList
//@description: 分页获取Branch_office记录
//@param: info request.Branch_officeSearch
//@return: err error, list interface{}, total int64

func GetBranch_officeInfoList(info request.Branch_officeSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Branch_office{})
    var branch_offices []model.Branch_office

    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&branch_offices).Error
	return err, branch_offices, total
}