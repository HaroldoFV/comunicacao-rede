# Comunicação em Rede - Cliente/Servidor TCP

Este projeto consiste em um simples programa cliente-servidor utilizando o protocolo TCP em Go. O servidor ecoa as mensagens recebidas do cliente de volta para ele.

## Estrutura do Projeto

O projeto contém dois arquivos principais:

1. `server.go`: Implementação do servidor TCP
2. `client.go`: Implementação do cliente TCP

## Servidor (server.go)

O servidor escuta na porta 9090 e aceita conexões de clientes. Ele lê as mensagens recebidas e as envia de volta ao cliente.

### Funcionalidades

- Escuta conexões TCP na porta 9090
- Aceita uma conexão de cliente
- Lê mensagens do cliente
- Ecoa as mensagens de volta para o cliente

## Cliente (client.go)

O cliente se conecta ao servidor na porta 9090 e permite que o usuário envie mensagens através do console. As respostas do servidor são exibidas na tela.

### Funcionalidades

- Conecta-se ao servidor TCP na porta 9090
- Lê entrada do usuário pelo console
- Envia mensagens para o servidor
- Recebe e exibe as respostas do servidor

## Como Executar

1. Inicie o servidor:
   ```
   go run server.go
   ```
2. Em outro terminal, inicie o cliente:
    ```
    go run client.go
    ```
3. No cliente, digite mensagens e pressione Enter. Você verá a resposta do servidor.

## Observações

- O servidor atual aceita apenas uma conexão por vez.
- O cliente continuará enviando mensagens até ser encerrado manualmente.
- Certifique-se de que a porta 9090 esteja disponível em sua máquina.

## Requisitos

- Go 1.x ou superior

## Possíveis Melhorias

- Implementar suporte a múltiplas conexões no servidor
- Adicionar tratamento de erros mais robusto
- Implementar um protocolo de comunicação mais sofisticado

## Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.