package proxy_item

import (
	"fmt"
	"strconv"
	"time"
)

// Proxy from https://www.sslproxies.org/
type ProxyItem struct {
	ProxyId   int
	ProxyIp   string
	ProxyPort int
	Country   Country
	Anonymity string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p ProxyItem) Key() string {
	return p.ProxyIp + ":" + strconv.Itoa(p.ProxyPort)
}

// NewProxyItem create ProxyItem from string values
func NewProxyItem(sIp string, sPort string, sCountryCode string, sCountryName string) (ProxyItem, error) {
	proxyPort, err := strconv.Atoi(sPort)
	if err != nil {
		return ProxyItem{}, fmt.Errorf("port should be int, got %s", sPort)
	}
	//proxyIp := net.ParseIP(sIp).To4()
	//
	//if proxyIp == nil {
	//	return ProxyItem{}, fmt.Errorf("ip should not be empty, got %s", sIp)
	//}
	return ProxyItem{
		ProxyIp:   sIp,
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
