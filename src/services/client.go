package services

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

// Client without timeout
func NewClient() *resty.Client {
	tr := &http.Transport{TLSHandshakeTimeout: 0 * time.Second}
	return resty.New().SetTransport(tr).EnableTrace()
}

func printTrace(resp *resty.Response) {
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("DNSLookup    :", ti.DNSLookup)
	fmt.Println("ConnTime     :", ti.ConnTime)
	fmt.Println("TLSHandshake :", ti.TLSHandshake)
	fmt.Println("ServerTime   :", ti.ServerTime)
	fmt.Println("ResponseTime :", ti.ResponseTime)
	fmt.Println("TotalTime    :", ti.TotalTime)
	fmt.Println("IsConnReused :", ti.IsConnReused)
	fmt.Println("IsConnWasIdle:", ti.IsConnWasIdle)
	fmt.Println("ConnIdleTime :", ti.ConnIdleTime)
}