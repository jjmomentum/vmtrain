package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/vmtrain/queue-monitor/pkg/common"
)

var (
	Type        = flag.String("Type", "simple", "The consumer. You can also set the KAFKA_PEERS environment variable")
	brokerList  = flag.String("brokers", os.Getenv("KAFKA_PEERS"), "The comma separated list of brokers in the Kafka cluster. You can also set the KAFKA_PEERS environment variable")
	dataManager = flag.String("datamanager", "", "The comma separated list of data managers . You can also set the KAFKA_PEERS environment variable")
	topic       = flag.String("topic", "", "REQUIRED: the topic to produce to")
	key         = flag.String("key", "", "The key of the message to produce. Can be empty.")
	value       = flag.String("value", "", "REQUIRED: the value of the message to produce. You can also provide the value on stdin.")
	partitioner = flag.String("partitioner", "", "The partitioning scheme to use. Can be `hash`, `manual`, or `random`")
	partition   = flag.Int("partition", -1, "The partition to produce to.")
	partitions  = flag.String("partitions", "all", "The partitions to consume, can be 'all' or comma-separated numbers")
	offset      = flag.String("offset", "newest", "The offset to start with. Can be `oldest`, `newest`")
	bufferSize  = flag.Int("buffer-size", 256, "The buffer size of the message channel.")
	verbose     = flag.Bool("verbose", false, "Turn on sarama logging to stderr")
	showMetrics = flag.Bool("metrics", false, "Output metrics on successful publish to stderr")
	silent      = flag.Bool("silent", false, "Turn off printing the message's topic, partition, and offset to stdout")
	action      = flag.String("action", "consume", "produce or consume")

	logger = log.New(os.Stderr, "", log.LstdFlags)
)

func main() {
	fmt.Printf("lab-reservation")

	flag.Parse()

	if *brokerList == "" {
		printUsageErrorAndExit("no -brokers specified. Alternatively, set the KAFKA_PEERS environment variable")
	}

	if *topic == "" {
		printUsageErrorAndExit("no -topic specified")
	}

	if *dataManager == "" {
		printUsageErrorAndExit("no -topic specified")
	}

	fmt.Printf("action: %s\n", *action)

	if *action == "produce" {
		err := common.ProduceMessage(*brokerList, *key, *topic, *value, int32(*partition), *partitioner)
		if err != nil {
			fmt.Errorf("error producing message")
			os.Exit(-1)
		}
	} else if *action == "consume" {

		if *Type == "simple" {
			sc := &common.SimpleConsumer{DataManager: *dataManager}
			err := sc.ConsumeMessages(*brokerList, *topic)
			if err != nil {
				fmt.Errorf("error")
				os.Exit(1)
			}
			return
		}

		err := common.ConsumeMessages(*brokerList, *topic, *bufferSize, *offset, *partitions)
		if err != nil {
			fmt.Errorf("error consuming message")
			os.Exit(-1)
		}
	}

}

func printErrorAndExit(code int, format string, values ...interface{}) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", fmt.Sprintf(format, values...))
	fmt.Fprintln(os.Stderr)
	os.Exit(code)
}

func printUsageErrorAndExit(message string) {
	fmt.Fprintln(os.Stderr, "ERROR:", message)
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Available command line options:")
	flag.PrintDefaults()
	os.Exit(64)
}

/*func stdinAvailable() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
*/
