package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/panupong-develop/go-simple/internal/ping"
)

func PingHandler(ctx *fiber.Ctx) error {
	usecases := ping.NewPingService()
	response := usecases.Pong()
	return NewResponse(ctx).ResponseOK(response, nil)
}
