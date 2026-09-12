package main

import (
	"bufio"
	"net"
	"os"
	"strconv"

	"github.com/NikosGour/logging/log"
	"github.com/NikosGour/mail_library/internal"
)

type Address struct {
	Domain string
	Port   uint16
}

func (addr Address) String() string {
	return string(addr.Domain) + ":" + strconv.Itoa(int(addr.Port))
}

func main() {
	addr := Address{Domain: "localhost", Port: 1025}
	log.Debug("Connecting to: `%s`", addr)
	conn, err := net.Dial("tcp", addr.String())
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

	err = internal.RunHelloCommand(conn)
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

	err = internal.RunMailCommand(conn)
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
