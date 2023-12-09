package main

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"hotel_train_antonyGG/types"
)

const KafkaTopic = "obudata"

type DataProducer interface {
	ProduceData(data types.OBUData) error
}

type KafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer() (DataProducer, error) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
	})

	if err != nil {
		return nil, err
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					//fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					//fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}

			}
		}
	}()

	if err != nil {
		return nil, err
	}

	return &KafkaProducer{
		producer: p,
	}, nil
}

func (p *KafkaProducer) ProduceData(data types.OBUData) error {

	b, err := json.Marshal(data)

	if err != nil {
		return err
	}

	topic := KafkaTopic

	return p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: b,
	}, nil)

}
