package request

import "gin-vue-admin/model"

type Wx_userSearch struct{
    model.Wx_user
    PageInfo
}