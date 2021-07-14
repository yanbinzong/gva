import service from '@/utils/request'

// @Tags user_paper_answer
// @Summary 创建user_paper_answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.user_paper_answer true "创建user_paper_answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user_paper_answer/createuser_paper_answer [post]
export const createuser_paper_answer = (data) => {
     return service({
         url: "/user_paper_answer/createUser_paper_answer",
         method: 'post',
         data
     })
 }


// @Tags user_paper_answer
// @Summary 删除user_paper_answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.user_paper_answer true "删除user_paper_answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user_paper_answer/deleteuser_paper_answer [delete]
 export const deleteuser_paper_answer = (data) => {
     return service({
         url: "/user_paper_answer/deleteUser_paper_answer",
         method: 'delete',
         data
     })
 }

// @Tags user_paper_answer
// @Summary 删除user_paper_answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除user_paper_answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user_paper_answer/deleteuser_paper_answer [delete]
 export const deleteuser_paper_answerByIds = (data) => {
     return service({
         url: "/user_paper_answer/deleteUser_paper_answerByIds",
         method: 'delete',
         data
     })
 }

// @Tags user_paper_answer
// @Summary 更新user_paper_answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.user_paper_answer true "更新user_paper_answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /user_paper_answer/updateuser_paper_answer [put]
 export const updateuser_paper_answer = (data) => {
     return service({
         url: "/user_paper_answer/updateUser_paper_answer",
         method: 'put',
         data
     })
 }


// @Tags user_paper_answer
// @Summary 用id查询user_paper_answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.user_paper_answer true "用id查询user_paper_answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /user_paper_answer/finduser_paper_answer [get]
 export const finduser_paper_answer = (params) => {
     return service({
         url: "/user_paper_answer/findUser_paper_answer",
         method: 'get',
         params
     })
 }

// @Tags user_paper_answer
// @Summary 分页获取user_paper_answer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取user_paper_answer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user_paper_answer/getuser_paper_answerList [get]
 export const getuser_paper_answerList = (params) => {
     return service({
         url: "/user_paper_answer/getUser_paper_answerList",
         method: 'get',
         params
     })
 }