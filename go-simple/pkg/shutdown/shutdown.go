package shutdown

import (
	"log"
	"os"
	"os/signal"
)

type shutdownFn func() error

func HandleShutdownGracefully(c chan os.Signal, appShutdown shutdownFn) {
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		log.Println("gracefully shutting down...")
		_ = appShutdown()
	}()
}
