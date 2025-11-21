package repository

import "context"

type KafkaProducer interface {
	Publish(ctx context.Context, topic string, key []byte, value []byte) error
}
