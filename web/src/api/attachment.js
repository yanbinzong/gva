import service from '@/utils/request'

// @Tags Attachment
// @Summary 创建Attachment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Attachment true "创建Attachment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /attachment/createAttachment [post]
export const createAttachment = (data) => {
     return service({
         url: "/attachment/createAttachment",
         method: 'post',
         data
     })
 }


// @Tags Attachment
// @Summary 删除Attachment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Attachment true "删除Attachment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /attachment/deleteAttachment [delete]
 export const deleteAttachment = (data) => {
     return service({
         url: "/attachment/deleteAttachment",
         method: 'delete',
         data
     })
 }

// @Tags Attachment
// @Summary 删除Attachment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Attachment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /attachment/deleteAttachment [delete]
 export const deleteAttachmentByIds = (data) => {
     return service({
         url: "/attachment/deleteAttachmentByIds",
         method: 'delete',
         data
     })
 }

// @Tags Attachment
// @Summary 更新Attachment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Attachment true "更新Attachment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /attachment/updateAttachment [put]
 export const updateAttachment = (data) => {
     return service({
         url: "/attachment/updateAttachment",
         method: 'put',
         data
     })
 }


// @Tags Attachment
// @Summary 用id查询Attachment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Attachment true "用id查询Attachment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /attachment/findAttachment [get]
 export const findAttachment = (params) => {
     return service({
         url: "/attachment/findAttachment",
         method: 'get',
         params
     })
 }


// @Tags Attachment
// @Summary 分页获取Attachment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Attachment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /attachment/getAttachmentList [get]
 export const getAttachmentList = (params) => {
     return service({
         url: "/attachment/getAttachmentList",
         method: 'get',
         params
     })
 }