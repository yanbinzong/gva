import service from '@/utils/request'

// @Tags Answer_question
// @Summary 创建Answer_question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Answer_question true "创建Answer_question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /answer_question/createAnswer_question [post]
export const createAnswer_question = (data) => {
     return service({
         url: "/answer_question/createAnswer_question",
         method: 'post',
         data
     })
 }

// @Tags Answer_question
// @Summary 删除Answer_question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Answer_question true "删除Answer_question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /answer_question/deleteAnswer_question [delete]
 export const deleteAnswer_question = (data) => {
     return service({
         url: "/answer_question/deleteAnswer_question",
         method: 'delete',
         data
     })
 }

// @Tags Answer_question
// @Summary 删除Answer_question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Answer_question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /answer_question/deleteAnswer_question [delete]
 export const deleteAnswer_questionByIds = (data) => {
     return service({
         url: "/answer_question/deleteAnswer_questionByIds",
         method: 'delete',
         data
     })
 }

// @Tags Answer_question
// @Summary 更新Answer_question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Answer_question true "更新Answer_question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /answer_question/updateAnswer_question [put]
 export const updateAnswer_question = (data) => {
     return service({
         url: "/answer_question/updateAnswer_question",
         method: 'put',
         data
     })
 }


// @Tags Answer_question
// @Summary 用id查询Answer_question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Answer_question true "用id查询Answer_question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /answer_question/findAnswer_question [get]
 export const findAnswer_question = (params) => {
     return service({
         url: "/answer_question/findAnswer_question",
         method: 'get',
         params
     })
 }

// @Tags Answer_question
// @Summary 分页获取Answer_question列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Answer_question列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /answer_question/getAnswer_questionList [get]
 export const getAnswer_questionList = (params) => {
     return service({
         url: "/answer_question/getAnswer_questionList",
         method: 'get',
         params
     })
 }

// @Tags Answer_question
// @Summary 创建Answer_question
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Answer_question true "创建Answer_question"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router/paper_answer/createUser_paper_answer [post]
export const createAnswer_question_front = (data) => {
    console.log(data)
    return service({
        url: "/answer_question/createAnswer_question_front",
        method: 'post',
        data
    })
}
