package main

import (
	"mini-socmed/internal/dependency"
	"mini-socmed/internal/httpserver"
)

func main() {
	logger := dependency.NewLogger()

	config, err := dependency.NewConfig(logger)
	if err != nil {
		return
	}

	db := dependency.ConnectDB(*config, logger)

	rc := dependency.NewRedisClient(*config, logger)
	if rc == nil {
		return
	}

	httpserver.InitApp(db, rc, *config, logger)
}
