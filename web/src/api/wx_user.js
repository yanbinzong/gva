import service from '@/utils/request'

// @Tags Wx_user
// @Summary 创建Wx_user
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Wx_user true "创建Wx_user"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wx_user/createWx_user [post]
export const createWx_user = (data) => {
     return service({
         url: "/wx_user/createWx_user",
         method: 'post',
         data
     })
 }


// @Tags Wx_user
// @Summary 删除Wx_user
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Wx_user true "删除Wx_user"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wx_user/deleteWx_user [delete]
 export const deleteWx_user = (data) => {
     return service({
         url: "/wx_user/deleteWx_user",
         method: 'delete',
         data
     })
 }

// @Tags Wx_user
// @Summary 删除Wx_user
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Wx_user"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wx_user/deleteWx_user [delete]
 export const deleteWx_userByIds = (data) => {
     return service({
         url: "/wx_user/deleteWx_userByIds",
         method: 'delete',
         data
     })
 }

// @Tags Wx_user
// @Summary 更新Wx_user
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Wx_user true "更新Wx_user"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wx_user/updateWx_user [put]
 export const updateWx_user = (data) => {
     return service({
         url: "/wx_user/updateWx_user",
         method: 'put',
         data
     })
 }


// @Tags Wx_user
// @Summary 用id查询Wx_user
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Wx_user true "用id查询Wx_user"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wx_user/findWx_user [get]
 export const findWx_user = (params) => {
     return service({
         url: "/wx_user/findWx_user",
         method: 'get',
         params
     })
 }


// @Tags Wx_user
// @Summary 分页获取Wx_user列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Wx_user列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wx_user/getWx_userList [get]
 export const getWx_userList = (params) => {
     return service({
         url: "/wx_user/getWx_userList",
         method: 'get',
         params
     })
 }