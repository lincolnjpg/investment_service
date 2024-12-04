package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/lincolnjpg/investment_service/internal/infra"
	"github.com/lincolnjpg/investment_service/internal/ports"
	"github.com/rabbitmq/amqp091-go"
)

type rabbitMqConsumer struct {
	financeApiBaseUrl    string
	investmentRepository ports.InvestmentRepository
}

func (c rabbitMqConsumer) consume() {
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

	var m infra.Message

	var forever chan struct{}

	go func() {
		for message := range messages {
			err := json.Unmarshal(message.Body, &m)
			if err != nil {
				log.Println("could not read message from queue:", err)
				os.Exit(1)
			}

			url := fmt.Sprintf("%s/%s.SA", c.financeApiBaseUrl, m.Ticker)

			r, err := http.Get(url)
			if err != nil {
				log.Fatalln(err)
			}

			if r.StatusCode != http.StatusOK {
				fmt.Println(r.StatusCode)
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

func NewRabbitMqConsumer(financeApiBaseUrl string, investmentRepository ports.InvestmentRepository) *rabbitMqConsumer {
	return &rabbitMqConsumer{
		financeApiBaseUrl:    financeApiBaseUrl,
		investmentRepository: investmentRepository,
	}
}
