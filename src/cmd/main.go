package main

import (
	"fmt"
	"github.com/aibotsoft/proxy-parser/src/services"
)

func main() {
	//print("Hello World")
	client := services.NewClient()
	doc, err := services.LoadNewProxy(client)
	if err != nil {
		panic(err)
	}
	proxyList, err := services.ProcessProxyPage(doc)
	if err != nil {
		panic(err)
	}

	//fmt.Println("Trace Info:", resp.Request.TraceInfo())
	//doc := resp.RawBody()
	//tr, err := htmlquery.QueryAll(doc, "//tbody/tr")

	//fmt.Println(resp.RawBody())
	fmt.Println(proxyList)

}
