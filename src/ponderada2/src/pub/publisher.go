package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// Função para gerar dados aleatórios de sensores
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
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("publisher")

	client := MQTT.NewClient(opts)
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

		token := client.Publish("/sensors", 1, false, msg) // QoS 1
		token.Wait()

		fmt.Println("Published:", msg)
		time.Sleep(2 * time.Second)
	}
}
