package main

import (
	"fmt"
	"github.com/aibotsoft/proxy-parser/src/services"
	"time"
)

func main() {
	client := services.NewClient()
	for {
		doc, err := services.LoadNewProxy(client)
		if err != nil {
			panic(err)
		}
		proxyList, err := services.ProcessProxyPage(doc)
		if err != nil {
			panic(err)
		}
		fmt.Println(proxyList)
		time.Sleep(1 * time.Minute)
	}
}

//fmt.Println("Trace Info:", resp.Request.TraceInfo())
//doc := resp.RawBody()
//tr, err := htmlquery.QueryAll(doc, "//tbody/tr")

//fmt.Println(resp.RawBody())
//fmt.Println(resp.Body())
//fmt.Printf("%T", rawBody)
//fmt.Println("asdfsadfasdf", tr)
//fmt.Printf("%v:\n", p)
//fmt.Printf("%v:\n", htmlquery.InnerText(td[0]))
