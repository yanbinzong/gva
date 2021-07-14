package request

import "gin-vue-admin/model"

type AttachmentSearch struct{
    model.Attachment
    PageInfo
}