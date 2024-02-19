---
sidebar_position: 1
---

# Simulador IoT

## Simulação de Envio e Captação de Dados de Sensores
Criar um simulador de dispositivos IoT utilizando o protocolo MQTT através do uso da biblioteca Eclipse Paho. 
Este projeto inclui uma estrutura de publisher e subscriber desenvolvida com Go, que publica mensagens JSON em um tópico MQTT e um que recebe, utilizando o Mosquitto como broker. As mensagens contêm dados simulados de medições, respeitando especificações técnicas como faixa de medição e alcance espectral, alem disso para garantir uma variação de dados, os valores são aleatórios dentro de um intervalo estabelecido pelo código.


### Repositório de Resolução do Projeto

[✔] [Ponderada 1](https://github.com/gabInteli/M9-Inteli-Eng-Comp_Gabriela_Matias/tree/main/src/ponderada1)

## Requisitos
- Mosquitto
- Paho MQTT
- Python
- Go

##  Modo de Execução 

### Configuração do Mosquitto
Você deve configurar o Mosquitto antes de iniciá-lo. Supondo que você tenha um arquivo de configuração chamado mosquitto.conf, você pode iniciar o Mosquitto com a seguinte linha de comando:

```
mosquitto -c mosquitto.conf
```

Esse comando inicia o broker Mosquitto com as configurações definidas em mosquitto.conf.

### Instalando Dependências - Go Mod
Acesse o diretorio que contem as dependências necessárias para cada função: 

Para o publisher:
```
/src/ponderada1/go/pub
```

Para o subscriber:
```
/src/ponderada1/go/sub
```

Acione as dependências para cada uma das pastas, com: 
```
go mod tidy
```

### Executando o Publisher

Para executar o script do publisher, navegue até o diretório onde o arquivo publisher.go: 
```
/src/ponderada1/go/pub
```

e execute o seguinte comando no terminal:
```
go run publisher.go
```

Isso iniciará o script que publicará mensagens JSON simuladas para o tópico MQTT no intervalo definido. 
O script imprimirá no terminal cada mensagem que for publicada.

### Executando o Subscriber

Para executar o script do subscriber, navegue até o diretório onde o arquivo subscriber.go: 
```
/src/ponderada1/go/sub
```

e execute o seguinte comando no terminal:
```
go run subscriber.go
```

Isso iniciará o script que receberá as mensagens JSON simuladas para o tópico MQTT no intervalo definido. O script imprimirá no terminal cada mensagem que for recebida.



### Demonstração: 

A demonstração pode ser verificada no vídeo abaixo:  
<iframe width="560" height="315" src="https://www.youtube.com/embed/8cfrfNcGn1A?si=bEVPjYMh1F_axd_t" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>
