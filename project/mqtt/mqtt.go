package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main() {
	var broker = "127.0.0.1"
	var port = 11883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername("emqx")
	opts.SetPassword("public")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//sub(client)
	publish(client)

	client.Disconnect(250)
}

func publish(client mqtt.Client) {
	num := 100
	for i := 0; i < num; i++ {
		text := "{\"industry\":\"\",\"scenario\":\"tightening\",\"type\":\"tightening\",\"entity_id\":\"12B62230/0000031604/1682648349\",\"final_object\":{\"entity_id\":\"12B62230/0000031604/1682648349\",\"scenario\":\"tightening\",\"factory_code\":\"\",\"product\":{\"track_code\":\"DDD\",\"no\":\"T1-06M1_0_0_1\"},\"components\":[{\"track_code\":\"DDD\",\"no\":\"T1-06M1_0_0_1\",\"spec\":\"\"}],\"user_id\":\"0\",\"station_name\":\"T1-06M1\",\"equipment\":{\"name\":\"\",\"sn\":\"\",\"type\":\"\",\"tools\":[{\"sn\":\"12B62230\",\"unit\":\"T1-06M1\",\"channel\":0,\"type\":\"tightening_tool\"}]},\"program\":{\"type\":\"pset\",\"strategy\":\"ADW\",\"batch\":0,\"batch_count\":0},\"measure_result\":\"\",\"measure_id\":\"\",\"measure_time\":0,\"error_code\":\"\",\"control_time\":\"\",\"dimension\":\"\",\"tolerance\":\"\",\"object_steps\":\"\"},\"process_object\":{\"entity_id\":\"\",\"scenario\":\"\",\"serials\":{\"cur_m\":[0.027578028,0.03545975,0.031518888,0.031518888,0.031518888,0.031518888,0.031518888,0.027578028,0.031518888,0.03545975,0.03940061,0.03940061,0.03545975,0.03940061,0.043341473,0.03545975,0.03545975,0.031518888,0.027578028,0.023641167,0.027578028,0.023641167,0.023641167,0.027578028,0.03545975,0.03545975,0.03940061,0.03545975,0.03545975,0.031518888,0.027578028,0.031518888,0.043341473,0.051223196,0.043341473,0.043341473,0.051223196,0.05910092,0.055160057,0.055160057,0.055160057,0.043341473,0.055160057,0.055160057,0.055160057,0.051223196,0.051223196,0.05910092,0.05910092,0.051223196,0.03940061,0.047282334,0.047282334,0.03545975,0.03545975,0.03940061,0.043341473,0.031518888,0.031518888,0.055160057,0.055160057,0.055160057,0.047282334,0.047282334,0.051223196,0.055160057,0.055160057,0.06304178,0.047282334,0.023641167,0.03940061,0.03940061,0.027578028,0.015759444,0.007877721,0.011818583,0.03545975,0.051223196,0.047282334,0.051223196,0.06698264,0.06698264,0.06698264,0.055160057,0.03545975,0.047282334,0.043341473,0.023641167,0.027578028,0.023641167,0.05910092,0.05910092,0.06304178,0.0709235,0.08274209,0.09062381,0.0709235,0.047282334,0.055160057,0.051223196,0.023641167,0.027578028,0.043341473,0.043341473,0.051223196,0.06698264,0.07880122,0.08274209,0.027578028,0.043341473,0.023641167,0.007877721,0.015759444,0.047282334,0.03940061,0.074864365,0.086682945,0.08274209,0.03940061,0.051223196,0.03545975,0.019700306,0.05910092,0.06304178,0.08274209,0.10244639,0.10638325,0.074864365,0.05910092,0.055160057,0.015759444,0.051223196,0.055160057,0.05910092,0.07880122,0.08274209,0.03940061,0.03940061,0.011818583,0.027578028,0.03545975,0.047282334,0.08274209,0.086682945,0.055160057,0.051223196,0.015759444,0.043341473,0.05910092,0.07880122,0.12608756,0.11820584,0.0709235,0.06698264,0.011818583,0.047282334,0.03545975,0.086682945,0.09456467,0.03940061,0.03940061,0.031518888,0.03545975,0.043341473,0.09456467,0.10244639,0.051223196,0.051223196,0.0039368607,0.06304178,0.086682945,0.12608756,0.12608756,0.06698264,0.03940061,0.023641167,0.03940061,0.07880122,0.09850553,0.03940061,0.03545975,0.03940061,0.027578028,0.055160057,0.10638325,0.055160057,0.055160057,0.015759444,0.06304178,0.09062381,0.13396528,0.09456467,0.07880122,0.019700306,0.031518888,0.0709235,0.110324115,0.051223196,0.03940061,0.043341473,0.023641167,0.110324115,0.10638325,0.06304178,0.019700306,0.06304178,0.09062381,0.141847,0.09456467,0.07880122,0.019700306,0.03940061,0.10244639,0.09456467,0.031518888,0.047282334,0.019700306,0.10244639,0.110324115,0.05910092,0.019700306,0.05910092,0.122146696,0.14578786,0.09062381,0.03940061,0.03940061,0.074864365,0.110324115,0.03940061,0.047282334,0.027578028,0.10638325,0.110324115,0.05910092,0.019700306,0.055160057,0.14578786,0.13396528,0.07880122,0.023641167,0.031518888,0.10244639,0.043341473,0.05910092,0.031518888,0.10638325,0.122146696,0.06304178,0.043341473,0.0709235,0.14972873,0.09456467,0.0709235,0.03545975,0.09850553,0.11820584,0.027578028,0.055160057,0.051223196,0.12608756,0.06304178,0.027578028,0.06304178,0.14972873,0.10638325,0.051223196,0.019700306,0.11426497,0.09456467,0.06304178,0.051223196,0.11426497,0.122146696,0.06304178,0.055160057,0.122146696,0.1615473,0.07880122,0.027578028,0.043341473,0.11426497,0.031518888,0.06304178,0.031518888,0.122146696,0.0709235,0.031518888,0.047282334,0.15366958,0.10638325,0.06304178,0.031518888,0.10244639,0.09456467,0.055160057,0.055160057,0.10244639,0.10244639,0.0709235,0.05910092,0.12608756,0.16942903,0.0709235,0.031518888,0.031518888,0.10638325,0.03545975,0.05910092,0.027578028,0.12608756,0.051223196,0.027578028,0.06304178,0.15760645,0.11820584,0.05910092,0.03940061,0.11426497,0.09062381,0.055160057,0.047282334,0.11426497,0.122146696,0.055160057,0.055160057,0.11426497,0.15366958,0.074864365,0.031518888,0.043341473,0.10638325,0.18912934,0.26005685,0.26005685,0.06698264,0.08274209,0.074864365,0.074864365],\"cur_w\":null,\"cur_t\":null,\"cur_s\":null}}}"
		token := client.Publish("topic/test", 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}

func sub(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}

//package main
//
//import (
//	"fmt"
//	"time"
//
//	mqtt "github.com/eclipse/paho.mqtt.golang"
//)
//
//const (
//	address  = "tcp://127.0.0.1:11883"
//	userName = ""
//	password = ""
//	topic    = "test/topic1"
//	clientID = "test"
//)
//
//var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
//	fmt.Printf("TOPIC: %s\n", msg.Topic())
//	fmt.Printf("MSG: %s\n", msg.Payload())
//}
//
//func main() {
//	opts := mqtt.NewClientOptions().
//		AddBroker(address).
//		SetClientID(clientID).
//		SetUsername(userName).
//		SetPassword(password)
//
//	opts.SetKeepAlive(60 * time.Second)
//	// 设置消息回调处理函数
//	opts.SetDefaultPublishHandler(f)
//	opts.SetPingTimeout(1 * time.Second)
//
//	c := mqtt.NewClient(opts)
//	if token := c.Connect(); token.Wait() && token.Error() != nil {
//		panic(token.Error())
//	}
//
//	// 断开连接
//	defer c.Disconnect(250)
//
//	// 发布消息
//	for {
//		token := c.Publish(topic, 0, false, "小镇编码家")
//		token.Wait()
//
//		time.Sleep(1 * time.Second)
//	}
//}
