package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/panupong-develop/go-simple/pkg/http/response"
)

func NewResponse(ctx *fiber.Ctx) response.IResponseHandler {
	return response.NewFiberResponseHandler(ctx)
}
