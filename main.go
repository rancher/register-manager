package main

import (
	"flag"
	"fmt"
	"github.com/rancher/register-tool/core"
	"os"
)

var (
	token      = ""
	resolveURL = ""
	url        = ""
)

func main() {
	flag.StringVar(&token, "registration-token", "", "Registration Token")
	flag.StringVar(&resolveURL, "resolve-url", "", "Resolve Url")
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
	if resolveURL != "" {
		// resolve url
		err := core.ResolveURL(resolveURL)
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
