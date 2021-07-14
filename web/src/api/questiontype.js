import service from '@/utils/request'

// @Tags Question_type
// @Summary 创建Question_type
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question_type true "创建Question_type"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /question_type/createQuestion_type [post]
export const createQuestion_type = (data) => {
     return service({
         url: "/question_type/createQuestion_type",
         method: 'post',
         data
     })
 }


// @Tags Question_type
// @Summary 删除Question_type
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question_type true "删除Question_type"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /question_type/deleteQuestion_type [delete]
 export const deleteQuestion_type = (data) => {
     return service({
         url: "/question_type/deleteQuestion_type",
         method: 'delete',
         data
     })
 }

// @Tags Question_type
// @Summary 删除Question_type
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Question_type"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /question_type/deleteQuestion_type [delete]
 export const deleteQuestion_typeByIds = (data) => {
     return service({
         url: "/question_type/deleteQuestion_typeByIds",
         method: 'delete',
         data
     })
 }

// @Tags Question_type
// @Summary 更新Question_type
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question_type true "更新Question_type"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /question_type/updateQuestion_type [put]
 export const updateQuestion_type = (data) => {
     return service({
         url: "/question_type/updateQuestion_type",
         method: 'put',
         data
     })
 }


// @Tags Question_type
// @Summary 用id查询Question_type
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question_type true "用id查询Question_type"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /question_type/findQuestion_type [get]
 export const findQuestion_type = (params) => {
     return service({
         url: "/question_type/findQuestion_type",
         method: 'get',
         params
     })
 }


// @Tags Question_type
// @Summary 分页获取Question_type列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Question_type列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /question_type/getQuestion_typeList [get]
 export const getQuestion_typeList = (params) => {
     return service({
         url: "/question_type/getQuestion_typeList",
         method: 'get',
         params
     })
 }

 // question list all
 export const getQuestion_typeList_all = (params) => {
    return  service({
        url:"/question_type/getQuestion_typeList_all",
        method: 'get',
        params
    })
 }