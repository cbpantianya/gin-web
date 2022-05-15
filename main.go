package main

import (
	_ "gin-template/handler/cron/out"
	_ "gin-template/handler/request/oauth"
	_ "gin-template/handler/request/ping"
	"gin-template/server"
)

func main() {
	server.Init()

	server.Start()

}
