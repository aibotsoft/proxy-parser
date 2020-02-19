package services

import (
	"github.com/aibotsoft/proxy-parser/src/models"
	"github.com/antchfx/htmlquery"
	"github.com/go-resty/resty/v2"
	"golang.org/x/net/html"
	"time"
)

const (
	GetProxyUrl    = "https://www.sslproxies.org/"
	proxyTablePath = `//table[@id="proxylisttable"]/tbody/tr`
	proxyIP        = 0
	proxyPort      = 1
	proxyCode      = 2
	proxyCountry   = 3
	proxyAnonymity = 4
)

// LoadURL loads the HTML document from the specified URL.
func LoadNewProxy(client *resty.Client) (*html.Node, error) {
	resp, err := client.R().SetDoNotParseResponse(true).Get(GetProxyUrl)
	if err != nil {
		return nil, err
	}
	printTrace(resp)
	body := resp.RawBody()
	defer body.Close()
	return html.Parse(body)
}

func ProcessProxyPage(doc *html.Node) ([]models.ProxyItem, error) {
	tr, err := htmlquery.QueryAll(doc, proxyTablePath)
	if err != nil {
		panic(err)
	}
	var proxyList []models.ProxyItem

	for _, row := range tr {
		td, err := htmlquery.QueryAll(row, "/td")
		if err != nil {
			panic(err)
		}
		p := models.ProxyItem{
			Ip:        htmlquery.InnerText(td[proxyIP]),
			Port:      htmlquery.InnerText(td[proxyPort]),
			Code:      htmlquery.InnerText(td[proxyCode]),
			Country:   htmlquery.InnerText(td[proxyCountry]),
			Anonymity: htmlquery.InnerText(td[proxyAnonymity]),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		}
		proxyList = append(proxyList, p)

	}
	return proxyList, nil
	//fmt.Println(resp.Body())
	//fmt.Printf("%T", rawBody)
	//fmt.Println("asdfsadfasdf", tr)
	//fmt.Printf("%v:\n", p)
	//fmt.Printf("%v:\n", htmlquery.InnerText(td[0]))

}
