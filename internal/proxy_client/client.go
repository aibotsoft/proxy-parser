package proxy_client

import (
	"github.com/aibotsoft/proxy-parser/internal/config"
	"github.com/aibotsoft/proxy-parser/internal/proxy_item"
	"github.com/antchfx/htmlquery"
	"github.com/go-resty/resty/v2"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"time"
)

const (
	proxyTablePath = `//table[@id="proxylisttable"]/tbody/tr`
	proxyIP        = 0
	proxyPort      = 1
	proxyCode      = 2
	proxyCountry   = 3
	proxyAnonymity = 4
)

type ProxyClient struct {
	cfg    *config.Config
	log    *log.Logger
	client *resty.Client
}

func (p *ProxyClient) CollectProxy() ([]proxy_item.ProxyItem, error) {
	pageNode, err := p.getNewProxy()
	if err != nil {
		return nil, err
	}
	return p.scrapeProxy(pageNode)
}

func (p *ProxyClient) scrapeProxy(proxyPageNode *html.Node) ([]proxy_item.ProxyItem, error) {
	tr, err := htmlquery.QueryAll(proxyPageNode, proxyTablePath)
	if err != nil {
		return nil, err
	}
	var proxyList []proxy_item.ProxyItem
	for _, row := range tr {
		td, err := htmlquery.QueryAll(row, "/td")
		if err != nil {
			p.log.Println("Error QueryAll", err)
			continue
		}
		item := proxy_item.ProxyItem{
			Ip:        htmlquery.InnerText(td[proxyIP]),
			Port:      htmlquery.InnerText(td[proxyPort]),
			Code:      htmlquery.InnerText(td[proxyCode]),
			Country:   htmlquery.InnerText(td[proxyCountry]),
			Anonymity: htmlquery.InnerText(td[proxyAnonymity]),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		}
		proxyList = append(proxyList, item)
	}
	return proxyList, nil
}

func (p *ProxyClient) getNewProxy() (*html.Node, error) {
	resp, err := p.client.R().SetDoNotParseResponse(true).Get(p.cfg.ProxyClient.Url)
	if err != nil {
		return nil, err
	}
	if p.cfg.IsDev() {
		p.logTrace(resp)
	}

	body := resp.RawBody()
	defer body.Close()
	return html.Parse(body)
}

func (p *ProxyClient) logTrace(resp *resty.Response) {
	p.log.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	p.log.Println("DNSLookup    :", ti.DNSLookup)
	p.log.Println("ConnTime     :", ti.ConnTime)
	p.log.Println("TLSHandshake :", ti.TLSHandshake)
	p.log.Println("ServerTime   :", ti.ServerTime)
	p.log.Println("ResponseTime :", ti.ResponseTime)
	p.log.Println("TotalTime    :", ti.TotalTime)
	p.log.Println("IsConnReused :", ti.IsConnReused)
	p.log.Println("IsConnWasIdle:", ti.IsConnWasIdle)
	p.log.Println("ConnIdleTime :", ti.ConnIdleTime)
}

// Client without timeout
func NewProxyClient(cfg *config.Config, log *log.Logger) (*ProxyClient, error) {
	pc := ProxyClient{cfg: cfg, log: log}
	tr := &http.Transport{TLSHandshakeTimeout: 0 * time.Second}
	pc.client = resty.New().SetTransport(tr).EnableTrace()
	return &pc, nil
}
