import service from '@/utils/request'

// @Tags Branch_office
// @Summary 创建Branch_office
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Branch_office true "创建Branch_office"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /branch_office/createBranch_office [post]
export const createBranch_office = (data) => {
     return service({
         url: "/branch_office/createBranch_office",
         method: 'post',
         data
     })
 }


// @Tags Branch_office
// @Summary 删除Branch_office
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Branch_office true "删除Branch_office"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /branch_office/deleteBranch_office [delete]
 export const deleteBranch_office = (data) => {
     return service({
         url: "/branch_office/deleteBranch_office",
         method: 'delete',
         data
     })
 }

// @Tags Branch_office
// @Summary 删除Branch_office
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Branch_office"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /branch_office/deleteBranch_office [delete]
 export const deleteBranch_officeByIds = (data) => {
     return service({
         url: "/branch_office/deleteBranch_officeByIds",
         method: 'delete',
         data
     })
 }

// @Tags Branch_office
// @Summary 更新Branch_office
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Branch_office true "更新Branch_office"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /branch_office/updateBranch_office [put]
 export const updateBranch_office = (data) => {
     return service({
         url: "/branch_office/updateBranch_office",
         method: 'put',
         data
     })
 }


// @Tags Branch_office
// @Summary 用id查询Branch_office
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Branch_office true "用id查询Branch_office"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /branch_office/findBranch_office [get]
 export const findBranch_office = (params) => {
     return service({
         url: "/branch_office/findBranch_office",
         method: 'get',
         params
     })
 }


// @Tags Branch_office
// @Summary 分页获取Branch_office列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Branch_office列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /branch_office/getBranch_officeList [get]
 export const getBranch_officeList = (params) => {
     return service({
         url: "/branch_office/getBranch_officeList",
         method: 'get',
         params
     })
 }