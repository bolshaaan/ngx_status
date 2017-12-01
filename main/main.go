package main

import (
	"fmt"
	"net/http"

	"flag"
	"os"

	"github.com/bolshaaan/nginx_status"
)

var hostParam = flag.String("hostname", "", "by default hostname of server")

func main() {

	flag.Parse()

	if *hostParam == "" {
		var err error
		*hostParam, err = os.Hostname()
		if err != nil {
			panic(err)
		}
	}

	resp, err := http.Get(fmt.Sprintf("http://%s/status/format/json", *hostParam))
	if err != nil {
		panic(err)
	}

	nginx_status.PrintStatus(resp.Body)
}
