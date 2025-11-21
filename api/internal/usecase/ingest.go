package usecase

import (
	"context"
	"encoding/json"
	"payment-ecosystem-clean/api/internal/domain"
	"payment-ecosystem-clean/api/internal/repository"
	"time"

	"github.com/google/uuid"
)

type IngestService struct {
	producer repository.KafkaProducer
	topic    string
}

func NewIngestService(p repository.KafkaProducer, topic string) *IngestService {
	return &IngestService{producer: p, topic: topic}
}
func (s *IngestService) Ingest(ctx context.Context, p *domain.Payment) (string, error) {
	p.ID = uuid.New().String()
	p.CreatedAt = time.Now().UTC()

	payload, err := json.Marshal(p)

	if err != nil {
		return "", err
	}

	if err := s.producer.Publish(ctx, s.topic, []byte(p.ID), payload); err != nil {
		return "", err
	}
	return p.ID, nil
}
