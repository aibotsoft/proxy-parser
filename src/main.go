package main

import "github.com/aibotsoft/proxy-parser/src/proxy_client"

func main() {
	print("Hello World")
	client := proxy_client.NewClient()
	proxy_client.GetProxy(client)
}
