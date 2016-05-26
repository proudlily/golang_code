package main

import (
	"net/http"
	"project/configs"
	"project/logs"
	"project/routes"
)

func main() {
	addr, _ := configs.MainConfig.String("server", "addr")
	logs.Logger.Info("Start server at:%v", addr)
	err := http.ListenAndServe(addr, routes.NewMux())
	logs.Logger.Critical("Server err:%v", err)
}
