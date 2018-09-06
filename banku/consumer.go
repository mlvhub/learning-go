package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

func mainConsumer(partition int32) {
	kafka := newKafkaConsumer()
	defer kafka.Close()

	config := cluster.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	consumer, err := cluster.NewConsumer(brokers, "banku-consumer", topics, config)
	if err != nil {
		fmt.Printf("Kafka error: %s\n", err)
		os.Exit(-1)
	}

	go consumeEvents(consumer)

	fmt.Println("Press [enter] to exit consumer\n")
	bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Terminating...")
}
