package serialize

// ResponseResult response result
type ResponseResult struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// BuildResponse serialize response data
func BuildResponse(code int64, msg string, data interface{}) ResponseResult {

	return ResponseResult{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
