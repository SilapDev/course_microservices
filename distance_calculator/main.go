package main

import (
	"log"
)

func main() {

	svc := NewCalcService()

	kafkaConsumer, err := NewKafkaConsumer("obudata", svc)

	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()

}
