package model

type ExcelInfo struct {
	FileName string        `json:"fileName"` // 文件名
	InfoList []SysBaseMenu `json:"infoList"`
}

// @增加问题结构体
type WenTiExcelInfo struct {
	FileName string 		`json:"fileName"` // 文件名
	InfoList []Question 	`json:"infoList"`
}