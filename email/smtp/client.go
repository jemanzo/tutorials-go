package smtpmail

import (
	"fmt"
	"log"
	"net/smtp"

	emailconf "github.com/jemanzo/tutorials-go/email/config"
)

// https://pkg.go.dev/net/smtp

func SendEmail(emailFrom string, emailTo string, conf *emailconf.Protocol) {
	// Connect to the remote SMTP server.
	conn, err := smtp.Dial(conf.GetURL())
	if err != nil {
		log.Fatal(err)
	}

	auth := smtp.CRAMMD5Auth(conf.Username, conf.Password)
	if err := conn.Auth(auth); err != nil {
		log.Fatal(err)
	}

	// Set the sender and recipient first
	if err := conn.Mail(emailFrom); err != nil {
		log.Fatal(err)
	}
	if err := conn.Rcpt(emailTo); err != nil {
		log.Fatal(err)
	}

	// Send the email body.
	wc, err := conn.Data()
	if err != nil {
		log.Fatal(err)
	}

	from := fmt.Sprintf("Compras <%s>", emailFrom)
	subject := "Testing subject 2"
	body := "Testing body 2"

	// _, err = fmt.Fprintf(wc, "Subject: %s\r\n\r\n%s\r\n", subject, body)
	fmt.Fprintf(wc, "From: %s\r\n", from)
	fmt.Fprintf(wc, "To: %s\r\n", emailTo)
	fmt.Fprintf(wc, "Subject: %s\r\n", subject)
	fmt.Fprintf(wc, "\r\n%s\r\n", body)
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Send the QUIT command and close the connection.
	err = conn.Quit()
	if err != nil {
		log.Fatal(err)
	}
}
