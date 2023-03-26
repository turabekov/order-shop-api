package main

import (
	"app/api"
	"app/config"
	"app/storage/postgresql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	store, err := postgresql.NewConnectPostgresql(&cfg)
	if err != nil {
		log.Println("Error connect to postgresql: ", err.Error())
		return
	}

	r := gin.New()

	// call logger
	r.Use(gin.Recovery(), gin.Logger())

	api.NewApi(r, &cfg, store)

	fmt.Println("Server running on port", cfg.ServerHost+cfg.ServerPort)
	err = r.Run(cfg.ServerHost + cfg.ServerPort)
	if err != nil {
		log.Println("Error listening server: ", err.Error())
		return
	}
}
