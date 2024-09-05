package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	connection, err := net.Dial("tcp", "127.0.0.1:9090")

	if err != nil {
		log.Fatal("Erro ao conectar com o servidor: ", err.Error())
	}

	for {
		message, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		_, errWrite := connection.Write([]byte(message))
		if errWrite != nil {
			log.Fatal("Erro: ", errWrite.Error())
		}

		response, _ := bufio.NewReader(connection).ReadString('\n')
		fmt.Println("SERVER: ", response)
	}
}
