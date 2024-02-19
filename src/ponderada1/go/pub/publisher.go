package main

import (
	"encoding/json"
	"fmt"
	"time"
	"math/rand"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("publisher")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		data := map[string]interface{}{
			"CO2":  rand.Intn(100),
			"CO":   rand.Intn(1000),
			"NO2":  rand.Intn(500),
			"MP10": rand.Intn(200),
			"MP2,5": rand.Intn(200),
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error converting data to JSON", err)
			return
		}

		msg := string(jsonData) + time.Now().Format(time.RFC3339)

		token := client.Publish("/sensors", 0, false, msg)
		token.Wait()

		fmt.Println("Published:", msg)
		time.Sleep(2 * time.Second)
	}
}