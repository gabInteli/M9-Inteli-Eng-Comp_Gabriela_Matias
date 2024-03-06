package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Teste de conexão com o broker MQTT
func TestConnectToMQTTBroker(t *testing.T) {
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
		t.Errorf("Falha ao conectar ao broker MQTT: %v", token.Error())
	} else {
		t.Log("Conexão com o broker MQTT estabelecida com sucesso.")
	}
}

// Teste de validação dos dados gerados
func TestGenerateData(t *testing.T) {
	data := GenerateData()

	// Verifique se todos os campos esperados estão presentes nos dados gerados
	expectedFields := []string{"CO2", "CO", "NO2", "MP10", "MP2,5"}
	for _, field := range expectedFields {
		if _, ok := data[field]; !ok {
			t.Errorf("Campo esperado %s não encontrado nos dados gerados", field)
			return
		}
	}
	t.Log("Dados gerados validados com sucesso.")
}

// Teste de confirmação de recebimento das publicações
func TestPublishAndReceiveMessages(t *testing.T) {
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
		t.Fatalf("Falha ao conectar ao broker MQTT: %v", token.Error())
	}

	// Inicia um subscriber MQTT para receber as publicações
	received := make(chan bool)
	token := client.Subscribe("/sensors", 0, func(client mqtt.Client, msg mqtt.Message) {
		// Verifique se a mensagem recebida é válida
		var data map[string]int
		if err := json.Unmarshal(msg.Payload(), &data); err != nil {
			t.Errorf("Erro ao decodificar a mensagem JSON: %v", err)
			return
		}

		// Verifique se todos os campos esperados estão presentes na mensagem recebida
		expectedFields := []string{"CO2", "CO", "NO2", "MP10", "MP2,5"}
		for _, field := range expectedFields {
			if _, ok := data[field]; !ok {
				t.Errorf("Campo esperado %s não encontrado na mensagem recebida", field)
				return
			}
		}

		// Marque como recebido
		received <- true
	})
	if token.Wait() && token.Error() != nil {
		t.Fatalf("Falha ao se inscrever no tópico MQTT: %v", token.Error())
	}

	// Publica uma mensagem com dados gerados
	data := GenerateData()
	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Erro ao converter dados para JSON: %v", err)
	}

	msg := string(jsonData)
	token = client.Publish("/sensors", 0, false, msg)
	if token.Wait() && token.Error() != nil {
		t.Fatalf("Falha ao publicar mensagem MQTT: %v", token.Error())
	}

	// Aguarda a confirmação de recebimento
	select {
	case <-received:
		// Mensagem recebida com sucesso
		t.Log("Mensagem recebida com sucesso.")
	case <-time.After(5 * time.Second):
		// Timeout - nenhum sinal de recebimento
		t.Fatalf("Timeout: Nenhuma mensagem recebida após 5 segundos")
	}
}
