package internal

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/NikosGour/logging/log"
)

type SMTPCommand = string

const (
	COMMAND_EXTENDED_HELLO SMTPCommand = "EHLO"
	COMMAND_HELLO          SMTPCommand = "HELO"
	COMMAND_MAIL           SMTPCommand = "MAIL"
	COMMAND_RECIPIENT      SMTPCommand = "RCPT"
	COMMAND_RESET          SMTPCommand = "RSET"
	COMMAND_VERIFY         SMTPCommand = "VRFY"
	COMMAND_EXPAND         SMTPCommand = "EXPN"
	COMMAND_HELP           SMTPCommand = "HELP"
	COMMAND_NOOP           SMTPCommand = "NOOP"
	COMMAND_QUIT           SMTPCommand = "QUIT"
)

func ReadResponses(responses chan<- string, conn *bufio.Reader, error chan<- error) {
	for {
		response, err := conn.ReadString('\n')
		if err != nil {
			error <- fmt.Errorf("on readString: %w", err)
			close(error)
			close(responses)
		}

		responses <- response

		// SMTP multiline responses have '-' after the status code.
		// The final response has a space.
		if len(response) >= 4 && response[3] == ' ' {
			close(responses)
			close(error)
			return
		}
	}
}

func ResponseHelper(responses <-chan string, error <-chan error, onRes func(res string), onErr func(err error)) {

loop:
	for {
		select {
		case err, ok := <-error:
			{
				if !ok {
					break loop
				}
				onErr(err)
			}

		case res, ok := <-responses:
			{
				if !ok {
					break loop
				}
				onRes(res)
			}
		}
	}

}

func RunSMTPCommand(conn io.ReadWriter, command SMTPCommand, args ...string) error {
	command_to_run := string(command + " " + strings.Join(args, " ") + "\r\n")
	log.Debug("sent command: `%s`", Unescape(command_to_run))

	n, err := conn.Write([]byte(command_to_run))
	if err != nil {
		return fmt.Errorf("on write: %w", err)
	}

	if n <= 0 {
		return fmt.Errorf("on write: no bytes writen")
	}

	return nil
}

func RunHelloCommand(conn io.ReadWriter) error {
	return RunSMTPCommand(conn, COMMAND_HELLO)
}

func RunExtendedHelloCommand(conn io.ReadWriter) error {
	return RunSMTPCommand(conn, COMMAND_EXTENDED_HELLO)
}

func RunNoopCommand(conn io.ReadWriter) error {
	return RunSMTPCommand(conn, COMMAND_NOOP)
}

func RunQuitCommand(conn io.ReadWriter) error {
	return RunSMTPCommand(conn, COMMAND_QUIT)
}

func RunHelpCommand(conn io.ReadWriter) error {
	return RunSMTPCommand(conn, COMMAND_HELP)
}

func RunMailCommand(conn io.ReadWriter) error {
	return RunSMTPCommand(conn, COMMAND_MAIL)
}

func RunRecipientCommand(conn io.ReadWriter) error {
	return RunSMTPCommand(conn, COMMAND_RECIPIENT)
}

func RunExpandCommand(conn io.ReadWriter) error {
	return RunSMTPCommand(conn, COMMAND_EXPAND)
}

func RunResetCommand(conn io.ReadWriter) error {
	return RunSMTPCommand(conn, COMMAND_RESET)
}

func RunVerifyCommand(conn io.ReadWriter) error {
	return RunSMTPCommand(conn, COMMAND_VERIFY)
}
