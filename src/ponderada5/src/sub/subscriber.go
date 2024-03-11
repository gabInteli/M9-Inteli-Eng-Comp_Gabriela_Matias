package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}

func main() {
	var broker = "bfd140d734d648f8858f07890dde8ff0.s1.eu.hivemq.cloud" // Endereço do broker HiveMQ na nuvem
	var port = 8883                                                    // Porta padrão para conexões não seguras

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d/mqtt", broker, port))
	opts.SetClientID("Subscriber")
	opts.SetUsername("admin")
	opts.SetPassword("Admin123")
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler


	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("/sensors", 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	fmt.Println("Subscriber está rodando. Pressione CTRL+C para sair.")
	select {} // Bloqueia indefinidamente
	

}