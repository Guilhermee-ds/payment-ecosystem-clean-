package main

import (
	"log"
	"os"

	delivery "payment-ecosystem-clean/api/internal/delivery/http"
	"payment-ecosystem-clean/api/internal/infra/kafka"
	"payment-ecosystem-clean/api/internal/usecase"
)

func main() {
	broker := getEnv("KAFKA_BROKER", "kafka:9092")
	topic := getEnv("KAFKA_TOPIC", "payments")
	addr := ":8080"

	// infra adapter
	producer := kafka.NewProducer(broker)
	defer producer.Close()

	// usecase
	ingestSvc := usecase.NewIngestService(producer, topic)

	// http handler
	h := delivery.NewHandler(ingestSvc)

	if err := delivery.Start(addr, h); err != nil {
		log.Fatalf("server crashed: %v", err)
	}
}

func getEnv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
