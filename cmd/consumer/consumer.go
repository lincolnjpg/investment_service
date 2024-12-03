package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Println("could not connect to RabbitMQ:", err)
		os.Exit(1)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Println("could not open a channel:", err)
		os.Exit(1)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("could not declare a queue:", err)
		os.Exit(1)
	}

	messages, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("could not register a consumer:", err)
		os.Exit(1)
	}

	var forever chan struct{}
	m := struct {
		AssetId uuid.UUID `json:"asset_id,omitempty"`
		Ticker  string    `json:"ticker,omitempty"`
	}{}

	go func() {
		for message := range messages {
			err := json.Unmarshal(message.Body, &m)
			if err != nil {
				log.Println("could not read message from queue:", err)
				os.Exit(1)
			}

			r, err := http.Get(fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s.SA", m.Ticker))
			if err != nil {
				log.Fatalln(err)
			}

			if r.StatusCode != http.StatusOK {
				log.Fatalln("Ticker not found")
			}

			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Fatalln(err)
			}

			var m map[string]any

			err = json.Unmarshal(body, &m)
			if err != nil {
				log.Fatalln(err)
			}

			chart, _ := m["chart"].(map[string]any)
			result, _ := chart["result"].([]any)
			resultZero, _ := result[0].(map[string]any)
			meta, _ := resultZero["meta"].(map[string]any)

			fmt.Println(meta["regularMarketPrice"])

			message.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
