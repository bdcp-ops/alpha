package main

import (
	"fmt"

	_ "github.com/bdcp-ops/alpha/aconfig"
	_ "github.com/bdcp-ops/alpha/aerror"
	_ "github.com/bdcp-ops/alpha/alog"
	_ "github.com/bdcp-ops/alpha/alog/gormwrapper"
	_ "github.com/bdcp-ops/alpha/autil"
	_ "github.com/bdcp-ops/alpha/autil/ahttp"
	_ "github.com/bdcp-ops/alpha/autil/ahttp/request"
	_ "github.com/bdcp-ops/alpha/database"
	_ "github.com/bdcp-ops/alpha/ginwrapper"
	_ "github.com/bdcp-ops/alpha/httpclient"
	_ "github.com/bdcp-ops/alpha/httpserver/rsp"
)

func main() {
	fmt.Println("Hello world")
}
