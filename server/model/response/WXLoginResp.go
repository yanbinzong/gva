package response
type WXloginResp struct {
	OpenId string 		`json:"openid"`
	SessionKey string	`json:"session_key"`
	UnionId string		`json:"union_id"`
	ErrCode int			`json:"err_code"`
	ErrMsg string		`json:"err_msg"`
}
