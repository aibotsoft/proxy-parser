package storage_test

import (
	"github.com/aibotsoft/proxy-parser/internal/config"
	"github.com/aibotsoft/proxy-parser/internal/logging"
	"github.com/aibotsoft/proxy-parser/internal/proxy_item"
	"github.com/aibotsoft/proxy-parser/internal/storage"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//func TestMain(m *testing.M)  {
//
//
//}
func TestNewStorage(t *testing.T) {
	s, err := initStorage(t)
	assert.Nil(t, err)
	assert.NotEmpty(t, s)
}

func initStorage(t *testing.T) (*storage.Storage, error) {
	t.Helper()
	cfg := config.NewConfig()
	log := logging.New(cfg)

	return storage.NewStorage(cfg, log)

}
func TestStorage_SaveProxy(t *testing.T) {
	s, _ := initStorage(t)

	testProxy := proxy_item.ProxyItem{
		Ip:   "1.1.1.1",
		Port: "80",
		Code: "US",
	}
	newProxy := s.SaveProxy(testProxy)
	assert.True(t, newProxy, "Add new proxy")
	time.Sleep(time.Millisecond)
	oldProxy := s.SaveProxy(testProxy)
	assert.False(t, oldProxy, "Add not new proxy")
}

func TestStorage_SaveProxyList(t *testing.T) {

}
