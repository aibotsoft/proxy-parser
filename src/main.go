package main

import (
	"github.com/aibotsoft/proxy-parser/src/proxy_client"
)

func main() {
	//print("Hello World")
	client := proxy_client.NewClient()
	doc, err := proxy_client.LoadNewProxy(client)
	if err != nil {
		panic("fuck off")
	}
	proxy_client.ProcessProxyPage(doc)
	//fmt.Println("Trace Info:", resp.Request.TraceInfo())
	//doc := resp.RawBody()
	//tr, err := htmlquery.QueryAll(doc, "//tbody/tr")

	//fmt.Println(resp.RawBody())
	//fmt.Println(resp.RawBody())

}
