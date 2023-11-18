package main

import (
	"mini-socmed/internal/dependency"
	"mini-socmed/internal/httpserver"
	"mini-socmed/internal/shared/helper"
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
	mdb, err := dependency.ConnectMongoDB(*config, logger)
	if err != nil {
		return
	}

	crypto := helper.NewAppCrypto()
	jwt := helper.NewJwtTokenizer()
	rsting := helper.NewRandomString()

	httpserver.InitApp(db, rc, mdb, *config, logger, crypto, jwt, rsting)
}
