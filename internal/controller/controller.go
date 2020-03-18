package controller

import (
	"github.com/aibotsoft/proxy-parser/internal/config"
	"github.com/aibotsoft/proxy-parser/internal/proxy_client"
	"github.com/aibotsoft/proxy-parser/internal/storage"
	"log"
	"time"
)

type Controller struct {
	cfg     *config.Config
	log     *log.Logger
	client  *proxy_client.ProxyClient
	storage *storage.Storage
}

func NewController(cfg *config.Config, log *log.Logger, client *proxy_client.ProxyClient) (*Controller, error) {
	c := &Controller{cfg: cfg, log: log, client: client}
	s, err := storage.NewStorage(cfg, log)
	if err != nil {
		return nil, err
	}
	c.storage = s
	return c, nil
}

func (c Controller) Job() {
	proxy, err := c.client.CollectProxy()
	if err != nil {
		c.log.Println("Error CollectProxy: ", err)
		return
	}
	c.storage.SaveProxyList(proxy)
}

func (c Controller) Run() error {
	c.Job()
	ticker := time.NewTicker(c.cfg.Controller.Period)
	defer ticker.Stop()
	for {
		select {
		case t := <-ticker.C:
			c.log.Println(t)
			c.Job()
		}
	}
}
