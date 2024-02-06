package service

import (
	"fmt"
	"time"

	"github.com/ShadamHarizky/publisher-service/messaging"
)

type PublisherService struct {
	publisher messaging.Publisher
}

func NewService(publisher messaging.Publisher) *PublisherService {
	return &PublisherService{
		publisher: publisher,
	}
}

func (s *PublisherService) Publish(key string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Publish ke publisher yang telah diinisialisasi di interface messaging
	err := s.publisher.Publish(key, timestamp)
	if err != nil {
		fmt.Printf("Error publishing: %v\n", err)
	}
}
