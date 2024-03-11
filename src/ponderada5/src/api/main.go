package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	//"log"
	"os"
	"os/signal"
	//"strconv"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/lib/pq"
)

type GasData struct {
	CO2       float64 `json:"CO2"`
	CO        float64 `json:"CO"`
	NO2       float64 `json:"NO2"`
	MP10      float64 `json:"MP10"`
	MP25      float64 `json:"MP25"`
	Timestamp string  `json:"timestamp"`
}

var mqttClient mqtt.Client

func startSubscriber() {
	var broker = "bfd140d734d648f8858f07890dde8ff0.s1.eu.hivemq.cloud" // Endereço do broker HiveMQ na nuvem
	var port = 8883                                                    // Porta padrão para conexões não seguras
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d/mqtt", broker, port))
	opts.SetClientID("Subscriber")
	opts.SetUsername("admin")
	opts.SetPassword("Admin123")
	opts.OnConnect = func(client mqtt.Client) {
		fmt.Println("Connected to MQTT broker")
		if token := client.Subscribe("/sensors", 1, onMessageReceived); token.Wait() && token.Error() != nil {
			fmt.Println("Error subscribing to /sensors", token.Error())
		}
	}
	mqttClient = mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	mqttClient.Disconnect(250)
	fmt.Println("Disconnected from MQTT broker")
}

func onMessageReceived(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	var gasData GasData
	err := json.Unmarshal(msg.Payload(), &gasData)
	if err != nil {
		fmt.Println("Error decoding JSON payload:", err)
		return
	}

	err = insertGasData(gasData)
	if err != nil {
		fmt.Println("Error inserting data into database:", err)
		return
	}

	fmt.Println("Data inserted into the database successfully.")
}

func insertGasData(gasData GasData) error {
	db, err := sql.Open("postgres", "postgresql://postgres:admin1234@127.0.0.53:5432/postgres?sslmode=disable")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Gas (co2, co, no2, mp10, mp25, timestamp) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		fmt.Println("Error preparing SQL statement:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(gasData.CO2, gasData.CO, gasData.NO2, gasData.MP10, gasData.MP25, gasData.Timestamp)
	if err != nil {
		fmt.Println("Error executing SQL statement:", err)
		return err
	}

	return nil
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}

func main() {
	startSubscriber()
}
