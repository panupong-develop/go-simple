package internal

import "github.com/panupong-develop/goapi/server/configs"

// Domain

type Pong struct {
	AppName string
	Version string
	Message string
}

// Application

type pingService struct {
}

func (h *pingService) Pong() *Pong {
	config := configs.GetConfig()
	return &Pong{
		AppName: config.App.Name,
		Version: config.App.Version,
		Message: "Pong!",
	}
}

func NewPingService() *pingService {
	return &pingService{}
}
