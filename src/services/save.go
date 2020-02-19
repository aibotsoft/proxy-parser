package services

import (
	"github.com/aibotsoft/proxy-parser/src/models"
	"github.com/go-resty/resty/v2"
)

// Send proxy list to proxy-backend
func SendProxyList(client *resty.Client, proxyList []models.ProxyItem) {
	//client.R().Post('')
}
