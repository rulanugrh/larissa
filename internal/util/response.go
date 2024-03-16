package util

type Response struct {
	Code    int         `json:"int"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Created(msg string, data interface{}) Response {
	return Response{
		Code:    201,
		Message: msg,
		Data:    data,
	}
}

func Success(msg string, data interface{}) Response {
	return Response{
		Code:    200,
		Message: msg,
		Data:    data,
	}
}

func NotFounds(msg string) Response {
	return Response{
		Code:    404,
		Message: msg,
	}
}

func InternalServerError(msg string) Response {
	return Response{
		Code:    500,
		Message: msg,
	}
}

func BadRequest(msg string) Response {
	return Response{
		Code:    400,
		Message: msg,
	}
}

func Unauthorized(msg string) Response {
	return Response{
		Code:    401,
		Message: msg,
	}
}

func Forbidden(msg string) Response {
	return Response{
		Code:    403,
		Message: msg,
	}
}

func Deleted(msg string) Response {
	return Response{
		Code:    204,
		Message: msg,
	}
}