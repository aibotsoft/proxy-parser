package main

import (
	"github.com/aibotsoft/proxy-parser/internal/config"
	"github.com/aibotsoft/proxy-parser/internal/controller"
	"github.com/aibotsoft/proxy-parser/internal/logging"
	"github.com/aibotsoft/proxy-parser/internal/proxy_client"
)

func main() {
	cfg := config.NewConfig()
	log := logging.New(cfg)
	log.Println("Beginning...")
	log.Printf("Config: %+v", cfg)

	client, err := proxy_client.NewProxyClient(cfg, log)
	if err != nil {
		log.Fatal(err)
	}
	c, err := controller.NewController(cfg, log, client)
	if err != nil {
		log.Fatal(err)
	}
	err = c.Run()
	if err != nil {
		log.Fatal(err)
	}
}
