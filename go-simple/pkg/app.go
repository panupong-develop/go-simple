package pkg

type App interface {
	LoadConfig()
	SetupDatabase()
	SetupMiddleware()
	SetupShutdownHandler()
	SetupRoutes()
	Start()
	Shutdown()
}
