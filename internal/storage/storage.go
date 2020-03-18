package storage

import (
	"github.com/aibotsoft/proxy-parser/internal/cache"
	"github.com/aibotsoft/proxy-parser/internal/config"
	"github.com/aibotsoft/proxy-parser/internal/proxy_item"
	"github.com/dgraph-io/ristretto"
	"github.com/nats-io/nats.go"
	"log"
)

type Storage struct {
	cfg   *config.Config
	log   *log.Logger
	cache *ristretto.Cache
	nc    *nats.Conn
}

func NewStorage(cfg *config.Config, log *log.Logger) (*Storage, error) {
	c, err := cache.NewCache()
	if err != nil {
		return nil, err
	}
	nc, err := Connect(cfg)
	if err != nil {
		return nil, err
	}
	return &Storage{cfg: cfg, log: log, cache: c, nc: nc}, nil
}

func (s Storage) SaveProxy(p proxy_item.ProxyItem) bool {
	//s.log.Println(p)
	proxyKey := p.Ip + ":" + p.Port
	//s.log.Println(proxyKey)

	_, ok := s.cache.Get(proxyKey)
	if ok {
		//s.log.Println("Proxy in cache: ")
		return false
	}
	//s.log.Println("TODO: Отправляем прокси в nc")
	//s.log.Println("Сохраняем в кеш")

	ok = s.cache.Set(proxyKey, true, 1)
	return true

}

func (s Storage) SaveProxyList(proxyList []proxy_item.ProxyItem) {
	countNewProxy := 0
	s.log.Printf("Получили %d прокси для проверки", len(proxyList))

	for _, p := range proxyList {
		newProxy := s.SaveProxy(p)
		if newProxy {
			countNewProxy++
			s.log.Printf("Добавили новое прокси %s", p.Ip)
		}
	}
	s.log.Printf("Всего добавили %d новых прокси", countNewProxy)
	s.log.Println(s.cache.Metrics)
}

func Connect(cfg *config.Config) (*nats.Conn, error) {
	natsConfig := cfg.Broker

	opts := nats.Options{
		Url:            natsConfig.Url,
		AllowReconnect: natsConfig.AllowReconnect,
		MaxReconnect:   natsConfig.MaxReconnect,
		ReconnectWait:  natsConfig.ReconnectWait,
		Timeout:        natsConfig.Timeout,
	}
	return opts.Connect()
}
