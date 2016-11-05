package main

import (
	"fmt"
	"flag"
	"os"
	"github.com/rancher/register-tool/core"
)

var (
	token = ""
	resolveUrl = ""
	url = ""
)

func main() {
	flag.StringVar(&token, "registration-token", "", "Registration Token")
	flag.StringVar(&resolveUrl, "resolve-url", "", "Resolve Url")
	flag.StringVar(&url, "load-url", "", "Load")
	flag.Parse()
	if token != "" {
		// register accessKey and secretKey
		err := core.Register(token)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	if resolveUrl != "" {
		// resolve url
		err := core.ResolveUrl(resolveUrl)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	if url != "" {
		// load env
		err := core.Load(url)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}