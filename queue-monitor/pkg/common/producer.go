package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Shopify/sarama"
)

func ProduceMessage(brokerList string, key string, topic string, value string, partition int32, partitioner string) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	switch partitioner {
	case "":
		if partition >= 0 {
			config.Producer.Partitioner = sarama.NewManualPartitioner
		} else {
			config.Producer.Partitioner = sarama.NewHashPartitioner
		}
	case "hash":
		config.Producer.Partitioner = sarama.NewHashPartitioner
	case "random":
		config.Producer.Partitioner = sarama.NewRandomPartitioner
	case "manual":
		config.Producer.Partitioner = sarama.NewManualPartitioner
		if partition == -1 {
			return fmt.Errorf("-partition is required when partitioning manually")
		}
	default:
		return fmt.Errorf(fmt.Sprintf("Partitioner %s not supported.", partitioner))
	}

	message := &sarama.ProducerMessage{Topic: topic, Partition: partition}

	if key != "" {
		message.Key = sarama.StringEncoder(key)
	}

	if value != "" {
		message.Value = sarama.StringEncoder(value)
	} else if stdinAvailable() {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("Failed to read data from the standard input: %s", err)
		}
		message.Value = sarama.ByteEncoder(bytes)
	} else {
		return fmt.Errorf("-value is required, or you have to provide the value on stdin")
	}

	producer, err := sarama.NewSyncProducer(strings.Split(brokerList, ","), config)
	if err != nil {
		fmt.Errorf("Failed to open Kafka producer: %s", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			logger.Println("Failed to close Kafka producer cleanly:", err)
		}
	}()

	partition, _, err = producer.SendMessage(message)
	if err != nil {
		return fmt.Errorf("Failed to produce message: %s", err)
	}
	//	if *showMetrics {
	//metrics.WriteOnce(config.MetricRegistry, os.Stderr)
	//}
	return nil
}

func stdinAvailable() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
