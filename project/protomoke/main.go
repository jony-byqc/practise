package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
)

type Service struct {
	ginHttp    *gin.Engine
	mqttClient mqtt.Client
}

func NewService() *Service {
	r := gin.Default()
	mqttCli := brokerConn()
	return &Service{
		ginHttp:    r,
		mqttClient: mqttCli,
	}
}

func main() {
	s := NewService()
	go s.httpRun()
	select {}

}
