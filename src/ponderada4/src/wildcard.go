package main

import (
	"fmt"
	"os"
	"time"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Recebido: %s de %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Conectado")
	client.Subscribe("sensor/+/temperatura", 1, nil)  // Usando o wildcard +
}

func main() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")
	opts.SetClientID("go_mqtt_client")
	opts.OnConnect = connectHandler
	opts.DefaultPublishHandler = messagePubHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	time.Sleep(60 * time.Second)
	client.Disconnect(250)
}