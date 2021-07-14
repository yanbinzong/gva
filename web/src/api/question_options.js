import service from '@/utils/request'

// @Tags Question_options
// @Summary 创建Question_options
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question_options true "创建Question_options"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /question_options/createQuestion_options [post]
export const createQuestion_options = (data) => {
     return service({
         url: "/question_options/createQuestion_options",
         method: 'post',
         data
     })
 }


// @Tags Question_options
// @Summary 删除Question_options
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question_options true "删除Question_options"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /question_options/deleteQuestion_options [delete]
 export const deleteQuestion_options = (data) => {
     return service({
         url: "/question_options/deleteQuestion_options",
         method: 'delete',
         data
     })
 }

// @Tags Question_options
// @Summary 删除Question_options
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Question_options"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /question_options/deleteQuestion_options [delete]
 export const deleteQuestion_optionsByIds = (data) => {
     return service({
         url: "/question_options/deleteQuestion_optionsByIds",
         method: 'delete',
         data
     })
 }

// @Tags Question_options
// @Summary 更新Question_options
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question_options true "更新Question_options"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /question_options/updateQuestion_options [put]
 export const updateQuestion_options = (data) => {
     return service({
         url: "/question_options/updateQuestion_options",
         method: 'put',
         data
     })
 }


// @Tags Question_options
// @Summary 用id查询Question_options
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Question_options true "用id查询Question_options"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /question_options/findQuestion_options [get]
 export const findQuestion_options = (params) => {
     return service({
         url: "/question_options/findQuestion_options",
         method: 'get',
         params
     })
 }


// @Tags Question_options
// @Summary 分页获取Question_options列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Question_options列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /question_options/getQuestion_optionsList [get]
 export const getQuestion_optionsList = (params) => {
     return service({
         url: "/question_options/getQuestion_optionsList",
         method: 'get',
         params
     })
 }