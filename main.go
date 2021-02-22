package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/vladbara705/NeuralNetwork/heplers"
	"log"
	"net/http"
	"time"
)

var cfg *ini.File

// Server type
type Server struct {
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
var ServerSetting = &Server{}

func setupServerConfig() {
	var err error
	cfg, err = ini.Load("server.ini")
	if err != nil {
		log.Fatalf("fail to parse 'server.ini': %v", err)
	}
	heplers.MapTo(cfg, "server", ServerSetting)
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

func main() {
	setupServerConfig()
	routesHandlers := setupRoutesConfig()
	endPoint := fmt.Sprintf(":%d", ServerSetting.HttpPort)
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routesHandlers,
		ReadTimeout:    ServerSetting.ReadTimeout,
		WriteTimeout:   ServerSetting.WriteTimeout,
	}

	log.Printf("[info] start http server listening %s", endPoint)
	_ = server.ListenAndServe()
}
