package response

type Response struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Payload interface{} `json:"payload"`
}

type options struct {
	success bool
	msg     string
	payload interface{}
}

type Option func(*options)

func SetSuccess(s bool) Option {
	return func(o *options) {
		o.success = s
	}
}

func SetMsg(msg string) Option {
	return func(o *options) {
		o.msg = msg
	}
}

func SetData(data interface{}) Option {
	return func(o *options) {
		o.payload = data
	}
}

func NewResponse(opts ...Option) Response {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	return Response{
		Success: o.success,
		Msg:     o.msg,
		Payload: o.payload,
	}
}

func NewResponseWithMsg(msg string) Response {
	return Response{
		Success: false,
		Msg:     msg,
		Payload: nil,
	}
}

func NewInvalidResponse() Response {
	return Response{
		Success: false,
		Msg:     MsgInvalidParam,
		Payload: nil,
	}
}

func NewInvalidResponseWithErr(err error) Response {
	return Response{
		Success: false,
		Msg:     err.Error(),
		Payload: nil,
	}
}

func NewErrorResponse() Response {
	return Response{
		Success: false,
		Msg:     MsgInternalServerError,
		Payload: nil,
	}
}

func NewSuccessResponse() Response {
	return Response{
		Success: true,
	}
}

func NewSuccessResponseWithPayload(payload interface{}) Response {
	return Response{
		Success: true,
		Payload: payload,
	}
}

func NewSuccessResponseWithMsg(msg string) Response {
	return Response{
		Success: true,
		Msg:     msg,
	}
}

func BadRequestResponse() Response {
	return Response{
		Success: false,
		Msg:     MsgInvalidParam,
	}
}
