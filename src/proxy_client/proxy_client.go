package proxy_client

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/go-resty/resty/v2"
	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
	"net/http"
	"time"
)

const (
	GetProxyUrl    = "https://www.sslproxies.org/"
	proxyTablePath = `//table[@id="proxylisttable"]/tbody/tr`
)

func NewClient() *resty.Client {
	tr := &http.Transport{TLSHandshakeTimeout: 0 * time.Second}
	client := resty.New().SetTransport(tr).EnableTrace()
	return client

}

// LoadURL loads the HTML document from the specified URL.
func LoadURL(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	return html.Parse(r)
}

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

func ProcessProxyPage(doc *html.Node) {
	tr, err := htmlquery.QueryAll(doc, proxyTablePath)
	if err != nil {
		panic("fuck off")
	}
	fmt.Printf("%v:\n", len(tr))

	for i, row := range tr {
		td, err := htmlquery.QueryAll(row, "/td")
		if err != nil {
			fmt.Printf("%v:\n", err)

			//panic(err)
		}
		if len(td) > 0 {
			//fmt.Printf("%v:\n %v", td, &row)
			fmt.Printf("%T:\n", td[0])
			fmt.Printf("%v:\n", htmlquery.InnerText(td[0]))
		} else {
			fmt.Print(i)

		}

	}

	//fmt.Println(resp.Body())
	//fmt.Printf("%T", rawBody)
	fmt.Println("asdfsadfasdf", tr)

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
