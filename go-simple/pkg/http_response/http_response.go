package http_response

// Implement Your Own Response Handler
// And use functions that we provide belows
type IResponseHandler interface {
	ResponseOK(
		data interface{},
		headers interface{},
	) error

	ResponseCreated(
		data interface{},
		headers interface{},
	) error

	ResponseBadRequest(
		errorCode string,
		errorMessage string,
		traceId string,
		data interface{},
	) error

	ResponseUnauthorized(
		errorCode string,
		errorMessage string,
		traceId string,
		data interface{},
	) error

	ResponsePermissionDenied(
		errorCode string,
		errorMessage string,
		traceId string,
		data interface{},
	) error

	ResponseInternalServerError(
		traceId string,
	) error
}

type Response struct {
	StatusCode int
	Headers    interface{}
	Success    *SuccessResponse
	Error      *ErrorResponse
}

type SuccessResponse struct {
	Data any
}

type ErrorResponse struct {
	ErrorCode    string
	ErrorMessage string
	TraceId      string
	Data         interface{}
}

func NewOKResponse(
	data interface{},
	headers interface{},
) *Response {
	return newSuccessResponse(200, data, headers)
}

func NewCreatedResponse(
	data interface{},
	headers interface{},
) *Response {
	return newSuccessResponse(201, data, headers)
}

func NewBadRequestResponse(
	errorCode string,
	errorMessage string,
	traceId string,
	data interface{},
) *Response {
	return newErrorResponse(400, errorCode, errorMessage, traceId, data)
}

func NewUnauthorizedResponse(
	errorCode string,
	errorMessage string,
	traceId string,
	data interface{},
) *Response {
	return newErrorResponse(401, errorCode, errorMessage, traceId, data)
}

func NewPermissionDeniedResponse(
	errorCode string,
	errorMessage string,
	traceId string,
	data interface{},
) *Response {
	return newErrorResponse(403, errorCode, errorMessage, traceId, data)
}

func NewInternalServerErrorResponse(
	traceId string,
) *Response {
	return newErrorResponse(500, "500", "Internal Server Error", traceId, nil)
}

func newSuccessResponse(
	statusCode int,
	data interface{},
	headers interface{},
) *Response {
	return &Response{
		StatusCode: statusCode,
		Headers:    headers,
		Success: &SuccessResponse{
			Data: data,
		},
		Error: nil,
	}
}

func newErrorResponse(
	statusCode int,
	errorCode string,
	errorMessage string,
	traceId string,
	data interface{},
) *Response {
	return &Response{
		StatusCode: statusCode,
		Headers:    nil,
		Success:    nil,
		Error: &ErrorResponse{
			ErrorCode:    errorCode,
			ErrorMessage: errorMessage,
			TraceId:      traceId,
			Data:         data,
		},
	}
}
