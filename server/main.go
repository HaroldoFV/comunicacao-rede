package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatal("Erro ao configurar o Listener.", err.Error())
	}

	connection, err := listener.Accept()
	if err != nil {
		log.Fatal("Erro ao iniciar uma nova conexão.", err.Error())
	}

	for {
		message, _ := bufio.NewReader(connection).ReadString('\n')
		_, err := connection.Write([]byte(message))
		if err != nil {
			log.Fatal("Erro: ", err.Error())
		}
	}
}
