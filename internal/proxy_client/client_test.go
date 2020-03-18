package proxy_client_test

import (
	"github.com/aibotsoft/proxy-parser/internal/config"
	"github.com/aibotsoft/proxy-parser/internal/logging"
	"github.com/aibotsoft/proxy-parser/internal/proxy_client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollectProxy(t *testing.T) {
	cfg := config.NewConfig()
	log := logging.New(cfg)
	client, err := proxy_client.NewProxyClient(cfg, log)
	assert.Nil(t, err)
	assert.NotEmpty(t, client)
	proxyList, err := client.CollectProxy()
	assert.Nil(t, err)
	assert.NotEmpty(t, proxyList)
	//t.Log(proxyList)
}

func TestCollectProxyTimeout(t *testing.T) {

}
