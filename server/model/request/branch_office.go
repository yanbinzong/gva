package request

import "gin-vue-admin/model"

type Branch_officeSearch struct{
    model.Branch_office
    PageInfo
}