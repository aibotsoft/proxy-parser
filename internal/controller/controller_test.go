package controller_test

import (
	"github.com/aibotsoft/proxy-parser/internal/config"
	"github.com/aibotsoft/proxy-parser/internal/controller"
	"github.com/aibotsoft/proxy-parser/internal/logging"
	"github.com/aibotsoft/proxy-parser/internal/proxy_client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewController(t *testing.T) {
	cfg := config.NewConfig()
	log := logging.New(cfg)
	client, err := proxy_client.NewProxyClient(cfg, log)
	assert.Nil(t, err)

	c, err := controller.NewController(cfg, log, client)
	assert.Nil(t, err)
	assert.NotEmpty(t, c)
	c.Job()
}
