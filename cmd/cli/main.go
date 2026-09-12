package main

import (
	"bufio"
	"net"
	"os"

	"github.com/NikosGour/logging/log"
	"github.com/NikosGour/mail_library/internal"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1025")
	if err != nil {
		log.Fatal("on dial tcp: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	responses := make(chan string)
	errCh := make(chan error)
	go internal.ReadResponses(responses, reader, errCh)
	internal.ResponseHelper(responses, errCh,
		func(res string) { log.Debug("SERVER: %s", res) },
		func(err error) { log.Fatal("on server read: %s", err) },
	)

	err = internal.RunExtendedHelloCommand(conn)
	if err != nil {
		log.Fatal("on RunHelloCommand: %v", err)
	}

	responses = make(chan string)
	errCh = make(chan error)
	go internal.ReadResponses(responses, reader, errCh)
	internal.ResponseHelper(responses, errCh,
		func(res string) { log.Debug("SERVER: %s", res) },
		func(err error) { log.Fatal("on server read: %s", err) },
	)

	err = internal.RunQuitCommand(conn)
	if err != nil {
		log.Fatal("on RunQuitCommand: %v", err)
	}

	responses = make(chan string)
	errCh = make(chan error)
	go internal.ReadResponses(responses, reader, errCh)
	internal.ResponseHelper(responses, errCh,
		func(res string) { log.Debug("SERVER: %s", res) },
		func(err error) { log.Fatal("on server read: %s", err) },
	)

}
