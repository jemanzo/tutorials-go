package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"syscall"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {
	// dir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(dir)
	// out, err := readFile("cat", "test.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var text = string(out)
	// regx := prepareRegX()
	// text2 := regx.FindAllString(text, -1)
	// for line := range text2 {
	// 	fmt.Println(text2[line])
	// }

	kafkaRun()

	// match, _ := regexp.MatchString("mar.a", string(out))
	// fmt.Printf("The date is %s\n", out)
}

func readFile(commandName, fileName string) ([]byte, error) {
	return exec.Command(commandName, fileName).Output()
}

func prepareRegX() *regexp.Regexp {
	r, err := regexp.Compile("--.*(Registered|Unregistered)")
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func kafkaRun() {

	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <broker> <group> <topics..>\n",
			os.Args[0])
		os.Exit(1)
	}

	broker := os.Args[1]
	group := os.Args[2]
	topics := os.Args[3:]

	fmt.Println("Broker:", broker)
	fmt.Println("Group:", group)
	fmt.Println("Topics:", topics)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		// Avoid connecting to IPv6 brokers:
		// This is needed for the ErrAllBrokersDown show-case below
		// when using localhost brokers on OSX, since the OSX resolver
		// will return the IPv6 addresses first.
		// You typically don't need to specify this configuration property.
		"broker.address.family": "v4",
		"group.id":              group,
		"session.timeout.ms":    6000,
		"auto.offset.reset":     "earliest"})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Consumer %v\n", c)

	err = c.SubscribeTopics(topics, nil)

	run := true

	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n",
					e.TopicPartition, string(e.Value))
				if e.Headers != nil {
					fmt.Printf("%% Headers: %v\n", e.Headers)
				}
			case kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				// But in this example we choose to terminate
				// the application if all brokers are down.
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	c.Close()
}
