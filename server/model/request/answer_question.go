package request

import "gin-vue-admin/model"

type Answer_questionSearch struct{
    model.Answer_question
    PageInfo
}