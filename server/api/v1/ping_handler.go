package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/panupong-develop/goapi/server/internal"
	pkg "github.com/panupong-develop/goapi/server/pkg/http_response"
)

func PingHandler(ctx *fiber.Ctx) error {
	usecases := internal.NewPingService()
	response := usecases.Pong()
	rHandler := pkg.NewFiberResponseHandler(ctx)
	return rHandler.ResponseOK(response, nil)
}
