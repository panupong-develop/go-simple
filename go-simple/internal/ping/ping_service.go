package ping

import "github.com/panupong-develop/go-simple/configs"

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
