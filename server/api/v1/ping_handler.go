package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/panupong-develop/go-simple/internal"
	pkg "github.com/panupong-develop/go-simple/pkg/http_response"
)

func PingHandler(ctx *fiber.Ctx) error {
	usecases := internal.NewPingService()
	response := usecases.Pong()
	rHandler := pkg.NewFiberResponseHandler(ctx)
	return rHandler.ResponseOK(response, nil)
}
