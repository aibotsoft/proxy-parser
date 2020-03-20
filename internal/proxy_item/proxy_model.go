package proxy_item

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// Proxy from https://www.sslproxies.org/
type ProxyItem struct {
	ProxyId   int
	ProxyIp   net.IP
	ProxyPort int
	Country   Country
	Anonymity string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p ProxyItem) Key() string {
	tcp := net.TCPAddr{
		IP:   p.ProxyIp,
		Port: p.ProxyPort,
	}
	return tcp.String()
}

// NewProxyItem create ProxyItem from string values
func NewProxyItem(sIp string, sPort string, sCountryCode string, sCountryName string) (ProxyItem, error) {
	proxyPort, err := strconv.Atoi(sPort)
	if err != nil {
		return ProxyItem{}, fmt.Errorf("port should be int, got %s", sPort)
	}
	proxyIp := net.ParseIP(sIp)
	if proxyIp == nil {
		return ProxyItem{}, fmt.Errorf("ip should not be empty, got %s", sIp)
	}
	return ProxyItem{
		ProxyIp:   proxyIp,
		ProxyPort: proxyPort,
		Country: Country{
			CountryName: sCountryName,
			CountryCode: sCountryCode,
			CreatedAt:   time.Now().UTC(),
		},
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}, nil
}

type Country struct {
	CountryId   int
	CountryName string
	CountryCode string
	CreatedAt   time.Time
}
