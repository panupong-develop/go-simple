package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/panupong-develop/go-simple/pkg/http_response"
)

func NewResponse(ctx *fiber.Ctx) http_response.IResponseHandler {
	return http_response.NewFiberResponseHandler(ctx)
}
