package response

import "gin-vue-admin/model"
type ExamFileResponse struct {
	File model.Attachment `json:"file"`
}