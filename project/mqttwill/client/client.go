package main

import (
	"fmt"
	"log"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// 创建MQTT客户端
	opts := MQTT.NewClientOptions().AddBroker("tcp://127.0.0.1:1883")
	opts.SetClientID("client1")

	// 设置第一个遗嘱消息
	willTopic1 := "topic/topic1"
	willPayload1 := "Will message 1"
	willQoS1 := 1
	opts.SetWill(willTopic1, willPayload1, byte(willQoS1), false)

	// 设置第二个遗嘱消息
	//willTopic2 := "topic/topic2"
	//willPayload2 := "Will message 2"
	//willQoS2 := 0
	//willRetain2 := true
	//opts.SetWill(willTopic2, willPayload2, byte(willQoS2), willRetain2)

	// 连接到MQTT代理服务器
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// 等待一段时间，以确保遗嘱消息已经发送
	time.Sleep(50 * time.Second)

	fmt.Println("Disconnected")
}
