package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

func CreateMessage(message model.Message) (err error) {
    //global.GVA_DB.AutoMigrate(&message)
	err = global.GVA_DB.Create(&message).Error
	return err
}