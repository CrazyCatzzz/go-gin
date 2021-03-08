package response

type Response struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Payload interface{} `json:"payload"`
}

func NewResponse(success bool, msg string, payload interface{}) Response {
	return Response{
		Success: success,
		Msg:     msg,
		Payload: payload,
	}
}

func NewErrorResponse(msg string) Response {
	return Response{
		Msg: msg,
	}
}

func NewSuccessResponse(msg string) Response {
	return Response{
		Success: true,
		Msg:     msg,
	}
}

func NewSuccessResponseWithPayload(payload interface{}, msg string) Response {
	return Response{
		Success: true,
		Msg:     msg,
		Payload: payload,
	}
}
