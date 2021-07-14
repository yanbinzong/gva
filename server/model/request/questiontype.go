package request

import "gin-vue-admin/model"

type Question_typeSearch struct{
    model.Question_type
    PageInfo
}