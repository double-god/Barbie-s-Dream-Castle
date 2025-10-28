package e
const(
	SUCCESS  =0
	ERROR    =1
	InvalidParams=400
	ErrorAuthCheckTokenFail=401
	ErrorAuthTokenTimeout=402
	ErrorAuthToken =403
	ErrorAuth =404

)
var MsgFlags = map[int]string{
	SUCCESS:                         "success",
	ERROR:                           "fail",
	InvalidParams:                   "请求参数错误",
	ErrorAuthCheckTokenFail:         "Token鉴权失败",
	ErrorAuthTokenTimeout:           "Token已超时",
	ErrorAuthToken:                  "Token生成失败",	
	ErrorAuth:                       "Token错误",
}
// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]		
}