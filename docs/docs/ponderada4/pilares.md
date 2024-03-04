---
sidebar_position: 2
---

# Pilares da CIA Triad 

### Confidencialidade:
Confidencialidade se refere à proteção de informações sensíveis contra acesso não autorizado. Isso significa garantir que apenas pessoas ou sistemas autorizados tenham permissão para acessar certos dados. Medidas de segurança para garantir a confidencialidade incluem criptografia de dados, controle de acesso baseado em permissões, autenticação forte e auditorias de acesso.

### Integridade:
Integridade diz respeito à garantia de que os dados permaneçam completos, precisos e confiáveis ao longo do tempo. Isso envolve prevenir, detectar e corrigir quaisquer modificações não autorizadas nos dados. Para manter a integridade dos dados, são utilizadas técnicas como assinaturas digitais, checksums, hash functions e sistemas de detecção de intrusões.

### Disponibilidade:
Disponibilidade refere-se à garantia de que os sistemas e recursos de informação estão sempre disponíveis quando necessário, sem interrupções ou atrasos indevidos. Isso implica em proteger os sistemas contra falhas, ataques de negação de serviço (DoS), e outros eventos que possam causar interrupções no serviço. Estratégias de alta disponibilidade, redundância de hardware e software, e planos de contingência são comumente utilizados para assegurar a disponibilidade dos sistemas.

## Resolução - Simulação/Descrição de Vulnerabilidades

### Violação do Pilar de Confidencialidade:
Uma possível violação da confidencialidade ocorreria se um atacante conseguisse interceptar o tráfego MQTT entre o cliente e o broker. Isso poderia ser feito através de ataques de sniffing na rede.
Outra forma de violação da confidencialidade seria se as credenciais de autenticação (username e password) fossem comprometidas, permitindo que um atacante não autorizado acessasse o sistema MQTT.

### Violação do Pilar de Integridade:
Uma violação da integridade poderia ocorrer se um atacante conseguisse modificar as mensagens MQTT em trânsito entre o cliente e o broker. Isso poderia resultar em informações incorretas sendo entregues aos destinatários.
Além disso, se um atacante conseguisse acesso não autorizado ao servidor MQTT, poderia modificar os tópicos ou mensagens armazenadas no broker, comprometendo a integridade dos dados.

### Violação do Pilar de Disponibilidade:
Uma forma de violação da disponibilidade seria um ataque de negação de serviço (DoS), onde um atacante tenta sobrecarregar o servidor MQTT com um grande volume de solicitações, tornando-o inacessível para usuários legítimos.
Além disso, se as configurações de recursos do contêiner forem muito restritivas, como no exemplo dado com limites muito baixos de CPU e memória, isso poderia resultar em indisponibilidade do serviço caso o servidor não seja capaz de lidar com a carga de trabalho.


### Exemplo de Simulação
Para simular essas violações, podemos: 

- Para a confidencialidade, capturar o tráfego MQTT usando ferramentas de sniffing de rede como o Wireshark.
- Para a integridade, podemos tentar modificar mensagens MQTT em trânsito ou alterar os dados armazenados no broker.
- Para a disponibilidade, podemos tentar sobrecarregar o servidor MQTT com um grande volume de solicitações, ou ajustar os recursos do contêiner para limitar severamente os recursos disponíveis e observar o impacto na disponibilidade do serviço.


