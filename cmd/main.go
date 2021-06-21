package main

import (
	"context"
	"crawler-backend/internal/crawler/infrastructure/router"
	"crawler-backend/internal/crawler/infrastructure/transport"
	"github.com/bearname/http-server/pkg/server"
	"github.com/bearname/url-extractor/pkg/app"
	"github.com/sirupsen/log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	crawler := app.New()
	controller := transport.New(crawler)
	handler := router.Router(controller)

	url := ":" + port
	log.WithFields(log.Fields{"url": url}).Info("starting the server")

	s := server.Server{}
	srv := s.StartServer(url, handler)

	s.WaitForKillSignal(s.GetKillSignalChan())
	err := srv.Shutdown(context.Background())
	if err != nil {
		log.Error(err)
		return
	}
}
