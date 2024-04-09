package main

import (
	"fmt"
	"time"
	"math/rand"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

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
	
	// Configurações do produtor
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  "pkc-7xoy1.eu-central-1.aws.confluent.cloud:9092",
		"sasl.mechanisms":    "PLAIN",
		"security.protocol":  "SASL_SSL",
		"sasl.username":      "NA7K2AO6VL7H22PI",
		"sasl.password":      "m+elheGPvsClhdIztpDa5k5ASAzR7v3Nm4syoU2G4uXCpCa8WG2cmiabuDEv6t0f",
	})
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	// Enviar mensagem
	topic := "sensor"

	for {
		data := GenerateData()
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error converting data to JSON", err)
			return
		}

	message := string(jsonData) 
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	producer.Flush(15 * 1000)
	}

	// Configurações do consumidor
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  "pkc-7xoy1.eu-central-1.aws.confluent.cloud:9092",
		"sasl.mechanisms":    "PLAIN",
		"security.protocol":  "SASL_SSL",
		"sasl.username":      "NA7K2AO6VL7H22PI",
		"sasl.password":      "m+elheGPvsClhdIztpDa5k5ASAzR7v3Nm4syoU2G4uXCpCa8WG2cmiabuDEv6t0f",
		"session.timeout.ms": 6000,
		"group.id":           "cluster_0",
		"auto.offset.reset":  "latest",  
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// Assinar tópico
	consumer.SubscribeTopics([]string{topic}, nil)

	// Consumir mensagens
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received message: %s\n", string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}
}