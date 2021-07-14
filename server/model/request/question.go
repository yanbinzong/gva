package request

import "gin-vue-admin/model"

type QuestionSearch struct{
    model.Question
    PageInfo
}