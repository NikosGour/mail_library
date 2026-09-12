package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/NikosGour/logging/log"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1025")
	if err != nil {
		log.Fatal("on dial tcp: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	res, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("on read: %v", err)
	}
	log.Debug("res: %v", res)

	n, err := conn.Write([]byte("EHLO localhost\r\n"))
	if err != nil {
		log.Fatal("on write: %v", err)
	}

	if n <= 0 {
		log.Fatal("no bytes read")
	}
	for {
		response, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		fmt.Print("SERVER: ", response)

		// SMTP multiline responses have '-' after the status code.
		// The final response has a space.
		if len(response) >= 4 && response[3] == ' ' {
			break
		}
	}

	// Quit
	fmt.Fprintf(conn, "QUIT\r\n")

	res, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Print("SERVER: ", res)
}
