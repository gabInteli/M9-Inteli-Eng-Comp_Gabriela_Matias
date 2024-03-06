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

func GenerateData() map[string]int {
	data := map[string]int{
		"CO2":   rand.Intn(100),
		"CO":    rand.Intn(1000),
		"NO2":   rand.Intn(500),
		"MP10":  rand.Intn(200),
		"MP2,5": rand.Intn(200),
	}
	return data
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}
	
	var broker = "bfd140d734d648f8858f07890dde8ff0.s1.eu.hivemq.cloud" // Endereço do broker HiveMQ na nuvem
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

		msg := string(jsonData) + time.Now().Format(time.RFC3339)
		token := client.Publish("/sensors", 1, false, msg)
		token.Wait()
		fmt.Println("Publicado:", msg)
		time.Sleep(2 * time.Second)
	}
}
