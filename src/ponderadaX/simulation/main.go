package main

import (
	"fmt"
	"time"
	"math/rand"
	"encoding/json"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}

func GenerateData() map[string]interface{} {
	data := map[string]interface{}{
		"idSensor":		rand.Intn(100),
		"tipoPoluente": rand.Intn(500),
		"nivel":        rand.Intn(1000),
		"Timestamp": 	time.Now().Format(time.RFC3339),
	}
	return data
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}
	
	var broker ="f553a9e9b8b54adab93346b920f9b07b.s1.eu.hivemq.cloud" // Endereço do broker HiveMQ na nuvem
	var port = 8883                                                    // Porta padrão para conexões não seguras

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d/mqtt", broker, port))
	opts.SetClientID("Publisher")
	opts.SetUsername("admin")
	opts.SetPassword("Admin123")
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		data := GenerateData()

		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error converting data to JSON", err)
			return
		}

		msg := string(jsonData)
		token := client.Publish("/sensor", 1, false, msg)
		token.Wait()
		fmt.Println("Publicado:", msg)
		time.Sleep(2 * time.Second)
	}
}