package response

import (
	"github.com/gofiber/fiber/v2"
)

type responseHandler struct {
	ctx *fiber.Ctx
}

func (h *responseHandler) asFiberResponse(r *Response) error {
	// TODO: Bind additional headers
	if r.StatusCode >= 400 {
		return h.ctx.Status(r.StatusCode).JSON(r.Error)
	} else {
		return h.ctx.Status(r.StatusCode).JSON(r.Success)
	}
}

func (r *responseHandler) ResponseOK(
	data interface{},
	headers interface{},
) error {
	return r.asFiberResponse(NewOKResponse(data, headers))
}

func (r *responseHandler) ResponseCreated(
	data interface{},
	headers interface{},
) error {
	return r.asFiberResponse(NewCreatedResponse(data, headers))
}

func (r *responseHandler) ResponseBadRequest(
	errorCode string,
	errorMessage string,
	traceId string,
	data interface{},
) error {
	return r.asFiberResponse(NewBadRequestResponse(errorCode, errorMessage, traceId, data))
}

func (r *responseHandler) ResponseUnauthorized(
	errorCode string,
	errorMessage string,
	traceId string,
	data interface{},
) error {
	return r.asFiberResponse(NewUnauthorizedResponse(errorCode, errorMessage, traceId, data))
}

func (r *responseHandler) ResponsePermissionDenied(
	errorCode string,
	errorMessage string,
	traceId string,
	data interface{},
) error {
	return r.asFiberResponse(NewPermissionDeniedResponse(errorCode, errorMessage, traceId, data))
}

func (r *responseHandler) ResponseInternalServerError(
	traceId string,
) error {
	return r.asFiberResponse(NewInternalServerErrorResponse(traceId))
}

func NewFiberResponseHandler(ctx *fiber.Ctx) IResponseHandler {
	return &responseHandler{ctx: ctx}
}
