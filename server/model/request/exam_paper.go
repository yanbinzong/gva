package request

import "gin-vue-admin/model"

type Exam_paperSearch struct{
    model.Exam_paper
    PageInfo
}