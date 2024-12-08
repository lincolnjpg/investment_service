package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/enum"
	"github.com/lincolnjpg/investment_service/internal/infra"
	"github.com/lincolnjpg/investment_service/internal/ports"
	"github.com/rabbitmq/amqp091-go"
)

type rabbitMqConsumer struct {
	financeApiBaseUrl    string
	assetRepository      ports.AssetRepository
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

			go c.processMessage(m)

			message.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func (c rabbitMqConsumer) processMessage(message infra.Message) {
	url := fmt.Sprintf("%s/%s.SA", c.financeApiBaseUrl, message.Ticker)

	r, err := http.Get(url)
	if err != nil {
		c.investmentRepository.UpdateInvestmentById(context.Background(), dtos.UpdateInvestmentByIdInput{Id: message.Investment.Id, Status: enum.Canceled})
		log.Fatalln(err)
	}

	if r.StatusCode != http.StatusOK {
		c.investmentRepository.UpdateInvestmentById(context.Background(), dtos.UpdateInvestmentByIdInput{Id: message.Investment.Id, Status: enum.Canceled})
		log.Fatalln("Ticker not found")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		c.investmentRepository.UpdateInvestmentById(context.Background(), dtos.UpdateInvestmentByIdInput{Id: message.Investment.Id, Status: enum.Canceled})
		log.Fatalln(err)
	}

	var response map[string]any

	err = json.Unmarshal(body, &response)
	if err != nil {
		c.investmentRepository.UpdateInvestmentById(context.Background(), dtos.UpdateInvestmentByIdInput{Id: message.Investment.Id, Status: enum.Canceled})
		log.Fatalln(err)
	}

	chart, _ := response["chart"].(map[string]any)
	result, _ := chart["result"].([]any)
	resultZero, _ := result[0].(map[string]any)
	meta, _ := resultZero["meta"].(map[string]any)
	tickerPrice, _ := meta["regularMarketPrice"].(float64)

	_, err = c.investmentRepository.UpdateInvestmentById(context.Background(), dtos.UpdateInvestmentByIdInput{Id: message.Investment.Id, Status: enum.Done})
	if err != nil {
		c.investmentRepository.UpdateInvestmentById(context.Background(), dtos.UpdateInvestmentByIdInput{Id: message.Investment.Id, Status: enum.Canceled})
		log.Fatalln(err)
	}

	_, err = c.assetRepository.UpdateAssetById(context.Background(), dtos.UpdateAssetByIdInput{
		Id:          message.Asset.Id,
		Name:        message.Asset.Name,
		UnitPrice:   tickerPrice,
		Rentability: message.Asset.Rentability,
		DueDate:     message.Asset.DueDate,
		Ticker:      message.Asset.Ticker,
		Type:        message.Asset.Type,
	})
	if err != nil {
		c.investmentRepository.UpdateInvestmentById(context.Background(), dtos.UpdateInvestmentByIdInput{Id: message.Asset.Id, Status: enum.Canceled})
		log.Fatalln(err)
	}
}

func NewRabbitMqConsumer(financeApiBaseUrl string, assetRepository ports.AssetRepository, investmentRepository ports.InvestmentRepository) *rabbitMqConsumer {
	return &rabbitMqConsumer{
		financeApiBaseUrl:    financeApiBaseUrl,
		assetRepository:      assetRepository,
		investmentRepository: investmentRepository,
	}
}
