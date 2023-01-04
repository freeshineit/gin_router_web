package models

// ResponseResult response result
type ResponseResult[T any] struct {
	// http status code
	Code int64 `json:"code"`
	// http message
	Msg string `json:"msg"`
	// http data
	Data T `json:"data"`
}

// BuildResponse serialize response data
func BuildResponse[T interface{}](code int64, msg string, data T) ResponseResult[T] {
	return ResponseResult[T]{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
