package request

import "gin-vue-admin/model"

type User_paper_answerSearch struct{
    model.User_paper_answer
    PageInfo
}