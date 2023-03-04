package helper

// ResponseResult response result
type ResponseResult[T any] struct {
	// http status code
	Code int64 `json:"code"`
	// http message
	Msg string `json:"msg,omitempty"`
	// http data
	Data T `json:"data,omitempty"`
}

// BuildResponse serialize response data
func BuildResponse[T interface{}](data T) ResponseResult[T] {
	return ResponseResult[T]{
		Code: 0,
		// Msg:  "success",
		Data: data,
	}
}

// BuildResponse serialize response data
func BuildErrorResponse[T interface{}](code int64, msg string) ResponseResult[T] {
	return ResponseResult[T]{
		Code: code,
		Msg:  msg,
	}
}
