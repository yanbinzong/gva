package request

import "gin-vue-admin/model"

type Question_optionsSearch struct{
    model.Question_options
    PageInfo
}