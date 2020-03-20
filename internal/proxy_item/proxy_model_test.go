package proxy_item

import (
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestNewProxyItem(t *testing.T) {
	sIp := "1.1.1.1"
	sPort := "80"
	sCountryCode := "US"
	sCountryName := "USA"

	t.Run("Valid proxy", func(t *testing.T) {
		pi, err := NewProxyItem(sIp, sPort, sCountryCode, sCountryName)
		assert.Nil(t, err)
		assert.Equal(t, net.ParseIP(sIp), pi.ProxyIp)
		assert.Equal(t, 80, pi.ProxyPort)
		assert.Equal(t, "US", pi.Country.CountryCode)
		assert.Equal(t, "USA", pi.Country.CountryName)
	})
	t.Run("Not_Valid port", func(t *testing.T) {
		_, err := NewProxyItem(sIp, "NotValidPort", sCountryCode, sCountryName)
		assert.Error(t, err)
	})
	t.Run("Empty ip address", func(t *testing.T) {
		_, err := NewProxyItem("", sPort, sCountryCode, sCountryName)
		assert.Error(t, err)
	})
	t.Run("Not_Valid ip address", func(t *testing.T) {
		_, err := NewProxyItem("NotValidIp", sPort, sCountryCode, sCountryName)
		assert.Error(t, err)
	})
}

func TestProxyItem_Key(t *testing.T) {
	sIp := "1.1.1.1"
	sPort := "80"
	pi, _ := NewProxyItem(sIp, sPort, "", "")
	got := pi.Key()
	want := "1.1.1.1:80"
	assert.Equal(t, want, got)
}
