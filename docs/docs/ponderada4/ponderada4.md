---
sidebar_position: 1
---

# Simulador IoT - Broker em Cloud

## Simulação de Envio e Captação de Dados de Sensores
Criar um simulador de dispositivos IoT utilizando o protocolo MQTT através do uso da biblioteca Eclipse Paho. 
Este projeto inclui uma estrutura de publisher e subscriber desenvolvida com Go, que publica mensagens JSON em um tópico MQTT e um que recebe, utilizando um broker hospedado em Cloud no HiveMQ. As mensagens contêm dados simulados de medições, respeitando especificações técnicas como faixa de medição e alcance espectral, alem disso para garantir uma variação de dados, os valores são aleatórios dentro de um intervalo estabelecido pelo código.


### Repositório de Resolução do Projeto

[✔] [Ponderada 4](https://github.com/gabInteli/M9-Inteli-Eng-Comp_Gabriela_Matias/tree/main/src/ponderada4)

## Requisitos
- Mosquitto
- Paho MQTT
- Python
- Go
- Cluster - HiveMQ  

##  Modo de Execução 

### Instalando Dependências - Go Mod
Acesse o diretório que contém as dependências necessárias para cada função: 

Para o publisher:
```
cd /src/ponderada4/src/pub
go mod tidy
```

Para o subscriber:
```
cd /src/ponderada4/src/sub
go mod tidy
```

### Executando o Publisher

Para executar o script do publisher, navegue até o diretório onde o arquivo publisher.go e execute o arquivo.
```
cd /src/ponderada2/src/pub
go run publisher.go

```

Isso iniciará o script que publicará mensagens JSON simuladas para o tópico MQTT no intervalo definido. 
O script imprimirá no terminal cada mensagem que for publicada.

![Publisher](../../static/img/pub_pond4.png)

### Executando o Subscriber

Para executar o script do subscriber, navegue até o diretório onde o arquivo subscriber.go e execute o arquivo: 
```
cd /src/ponderada2/src/pub
go run subscriber.go

```

Isso iniciará o script que receberá as mensagens JSON simuladas para o tópico MQTT no intervalo definido. O script imprimirá no terminal cada mensagem que for recebida.

![Subscriber](../../static/img/sub_pond4.png)

