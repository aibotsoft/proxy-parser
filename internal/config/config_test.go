package config_test

import (
	"github.com/aibotsoft/proxy-parser/internal/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestLoadEnv(t *testing.T) {
	err := config.LoadEnv()
	assert.Nil(t, err)
	assert.Equal(t, "true", os.Getenv("TEST_LOAD_ENV"))
}

func TestNew(t *testing.T) {
	cfg := config.NewConfig()
	assert.NotEmpty(t, cfg)
}

func TestService(t *testing.T) {
	cfg := config.NewConfig()
	assert.Equal(t, true, cfg.Service.TestLoadEnv)
	assert.Equal(t, "proxy-parser", cfg.Service.Name)

}
func TestProxyClient(t *testing.T) {
	cfg := config.NewConfig()
	assert.Equal(t, 5*time.Second, cfg.ProxyClient.Timeout)
	assert.Equal(t, "https://www.sslproxies.org/", cfg.ProxyClient.Url)
}
