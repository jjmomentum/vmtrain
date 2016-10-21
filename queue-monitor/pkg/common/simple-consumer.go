package common

import (
	//"encoding/base64"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	//"github.com/tdhite/q3-training-journal/journal"
)

type SimpleConsumer struct {
	DataManager string
}

func (sc *SimpleConsumer) saveReservation(msg *Message) {
	url := fmt.Sprintf("%s/api/reservations", sc.DataManager)
	fmt.Printf("Sending reservation %s to data-manager at %s\n", msg.ToJson(), url)
	payload := string(msg.Base64[:])
	fmt.Printf("payload: %s\n", payload)
	//var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	var jsonStr = []byte(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}

func (sc *SimpleConsumer) handleMessages(messages <-chan *Message, topic string) error {
	go func() {
		for msg := range messages {
			fmt.Printf("topic: %s", msg.ToJson())
			fmt.Println()
			if topic == "reservation" {
				sc.saveReservation(msg)
			}

		}
	}()

	return nil
}

func (sc *SimpleConsumer) consume(url string, topic string, listen chan bool) (<-chan *Message, error) {
	messages := make(chan *Message, 1)
	tl := fmt.Sprintf("%s/api/topic/%s", url, topic)

	go func() {
		for <-listen {
			fmt.Printf("Checking queue at %s for topic %s\n", url, topic)

			res, err := http.Get(tl)
			if err != nil {
				log.Print(err)
				continue
			}

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				//log.Print(err)
				continue
			}

			//fmt.Printf("raw: %s", b)
			msg := &Message{}

			msg.FromJson(b)
			if msg.Base64 == nil || len(msg.Base64) == 0 {
				//log.Print("no message")
				res.Body.Close()
				continue
			}
			messages <- msg
			res.Body.Close()
			if err != nil {
				log.Fatal(err)
			}

		}
		fmt.Println("Exit consumer")

		close(messages)
	}()

	return messages, nil
}

func (sc *SimpleConsumer) ConsumeMessages(url string, topic string) error {

	listen := make(chan bool, 1)

	messages, _ := sc.consume(url, topic, listen)
	err := sc.handleMessages(messages, topic)
	if err != nil {
		fmt.Errorf("%V", err)
		os.Exit(1)
	}

	for {
		listen <- true
		time.Sleep(time.Second * 15)
	}

	listen <- false
	time.Sleep(time.Second * 3)
	return nil

}
