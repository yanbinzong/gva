import service from '@/utils/request'

// @Tags Exam_paper
// @Summary 创建Exam_paper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Exam_paper true "创建Exam_paper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exam_paper/createExam_paper [post]
export const createExam_paper = (data) => {
     return service({
         url: "/exam_paper/createExam_paper",
         method: 'post',
         data
     })
 }


// @Tags Exam_paper
// @Summary 删除Exam_paper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Exam_paper true "删除Exam_paper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exam_paper/deleteExam_paper [delete]
 export const deleteExam_paper = (data) => {
     return service({
         url: "/exam_paper/deleteExam_paper",
         method: 'delete',
         data
     })
 }

// @Tags Exam_paper
// @Summary 删除Exam_paper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Exam_paper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exam_paper/deleteExam_paper [delete]
 export const deleteExam_paperByIds = (data) => {
     return service({
         url: "/exam_paper/deleteExam_paperByIds",
         method: 'delete',
         data
     })
 }

// @Tags Exam_paper
// @Summary 更新Exam_paper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Exam_paper true "更新Exam_paper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exam_paper/updateExam_paper [put]
 export const updateExam_paper = (data) => {
    console.log(data)
     return service({
         url: "/exam_paper/updateExam_paper",
         method: 'put',
         data
     })
 }


// @Tags Exam_paper
// @Summary 用id查询Exam_paper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Exam_paper true "用id查询Exam_paper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exam_paper/findExam_paper [get]
 export const findExam_paper = (params) => {
     return service({
         url: "/exam_paper/findExam_paper",
         method: 'get',
         params
     })
 }


// @Tags Exam_paper
// @Summary 分页获取Exam_paper列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Exam_paper列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exam_paper/getExam_paperList [get]
 export const getExam_paperList = (params) => {
     return service({
         url: "/exam_paper/getExam_paperList",
         method: 'get',
         params
     })
 }


// @Tags Exam_paper
// @Summary 获取试卷->全部问题
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Exam_paper列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exam_paper/findExam_paper_question [get]
// export const findExam_paperQuestion = (params) => {
//     return service({
//         url:"exam_paper/findExam_paperQuestion",
//         method: 'get',
//         params
//     })
// }
export const findExam_paperQuestion = (params) => {
    return service({
        url: "/exam_paper/findExam_paperQuestion",
        method: 'get',
        params
    })
}
