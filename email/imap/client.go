package imapmail

import (
	"fmt"
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	emailconf "github.com/jemanzo/tutorials-go/email/config"
)

type IMAPClient struct {
	conf emailconf.IMAP
	conn *client.Client
}

func (c *IMAPClient) SetConfig(conf emailconf.IMAP) {
	c.conf = conf
}

func (c *IMAPClient) Connect() {
	// Connect to server
	url := fmt.Sprintf("%s:%d", c.conf.Host, c.conf.Port)
	log.Printf("Connecting to server %s", url)
	conn, err := client.DialTLS(url, nil)
	if err != nil {
		log.Fatal(err)
	}
	c.conn = conn
	log.Println("Connected")
}

func (c *IMAPClient) Close() {
	c.conn.Logout()
}

func (c *IMAPClient) Login() {
	// Don't forget to logout by calling Close()
	if err := c.conn.Login(c.conf.Username, c.conf.Password); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")
}

func (c *IMAPClient) ListMailboxes() {
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.conn.List("", "*", mailboxes)
	}()

	log.Println("Mailboxes:")
	for m := range mailboxes {
		log.Println("* " + m.Name)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}
}

func (c *IMAPClient) SelectINBOX() (*imap.MailboxStatus, error) {
	return c.Select("INBOX")
}

func (c *IMAPClient) Select(folder string) (*imap.MailboxStatus, error) {
	mbox, err := c.conn.Select(folder, false)
	if err != nil {
		return nil, err
	}
	log.Printf("Flags for %s: %v\n", folder, mbox.Flags)
	return mbox, err
}

func (c *IMAPClient) Messages(mbox *imap.MailboxStatus) {
	// Get the last 4 messages
	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > 3 {
		// We're using unsigned integers here, only subtract if the result is > 0
		from = mbox.Messages - 3
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.conn.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()

	log.Println("Last 4 messages:")
	for msg := range messages {
		log.Println("* " + msg.Envelope.Subject)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}
	log.Println("Done!")
}
