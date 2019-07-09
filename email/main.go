package main

import (
	"fmt"
	"log"

	emailconf "github.com/jemanzo/tutorials-go/email/config"
	imapmail "github.com/jemanzo/tutorials-go/email/imap"
	smtpmail "github.com/jemanzo/tutorials-go/email/smtp"
)

const CFG_FILENAME = "./emails.yaml"
const CFG_FILENAME_SAMPLE = "./emails.sample.yaml"

func main() {
	// Create or Read Email Config
	conf, err := ReadConfig(CFG_FILENAME)
	if err != nil {
		conf, err = CreateConfigSample(CFG_FILENAME_SAMPLE)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(conf)

	for key := range conf.Emails {
		emailFrom := key
		emailConf := conf.Emails[key]
		log.Printf("found email: %s\n", emailFrom)
		log.Printf("testing SMTP for %s\n", emailFrom)
		smtpmail.SendEmail(emailFrom, emailFrom, emailConf.SMTP)
	}

	// runIMAP(email.IMAP)
}

func runIMAP(imapCfg *emailconf.Protocol) {
	imap := imapmail.IMAPClient{}
	imap.SetConfig(*imapCfg)
	imap.Connect()
	imap.Login()
	imap.ListMailboxes()
	mbox, _ := imap.SelectINBOX()
	imap.Messages(mbox)
}
