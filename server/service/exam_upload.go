package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils/upload"
	"mime/multipart"
	"strings"
)


//@function: Upload
//@description: 创建文件上传记录
//@param: file model.Attachment
//@return: error

func ExamUpload(file *model.Attachment) error {
	return global.GVA_DB.Create(&file).Error
}


//@function: FindFile
//@description: 删除文件切片记录
//@param: id uint
//@return: error, model.Attachment

func ExamFindFile(id uint) (error, model.Attachment) {
	var file model.Attachment
	err := global.GVA_DB.Where("id = ?", id).First(&file).Error
	return err, file
}


//@function: DeleteFile
//@description: 删除文件记录
//@param: file model.Attachment
//@return: err error

func ExamDeleteFile(file model.Attachment) (err error) {
	var fileFromDb model.Attachment
	err, fileFromDb = ExamFindFile(file.ID)
	oss := upload.NewOss()
	if err = oss.DeleteFile(fileFromDb.Key); err != nil {
		return errors.New("文件删除失败")
	}
	err = global.GVA_DB.Where("id = ?", file.ID).Unscoped().Delete(&file).Error
	return err
}


//@function: GetFileRecordInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func ExamGetFileRecordInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var fileLists []model.Attachment
	err = db.Find(&fileLists).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return err, fileLists, total
}


//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: err error, file model.Attachment

func ExamUploadFile(header *multipart.FileHeader, noSave string, typeString string) (err error, file model.Attachment) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := model.Attachment{
			Url:  filePath,
			Name: header.Filename,
			Exts:  s[len(s)-1],
			Key:  key,
			Type: typeString,
		}
		return ExamUpload(&f), f
	}
	return
}
