// main.go
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ShadamHarizky/publisher-service/messaging"
	"github.com/ShadamHarizky/publisher-service/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Inisialisasi koneksi
	var publisher messaging.Publisher
	var key string

	// Pilih jenis publisher dari env
	publisherType := os.Getenv("PUBLISHER_TYPE")
	switch publisherType {
	case "rabbitmq":
		rabbitMQURL := os.Getenv("RABBITMQ_URL")
		key = os.Getenv("RABBITMQ_ROUTING_KEY")
		rabbitMQPublisher, err := messaging.NewRabbitMQPublisher(rabbitMQURL)
		if err != nil {
			log.Fatalf("Failed to initialize RabbitMQ publisher: %v", err)
		}
		defer rabbitMQPublisher.Close()
		publisher = rabbitMQPublisher
	case "redis":
		redisPublisher := messaging.NewRedisPublisher(os.Getenv("REDIS_ADDRESS"))
		key = os.Getenv("REDIS_KEY")
		defer redisPublisher.Close()
		publisher = redisPublisher
	default:
		log.Fatalf("Invalid publisher type: %s", publisherType)
	}

	service := service.NewService(publisher)

	// Menjalankan publisher setiap detik
	go func() {
		for {
			service.Publish(key)
			// Tunggu satu detik sebelum mengirim pesan lagi
			time.Sleep(time.Second)
		}
	}()

	// close connection
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown
	os.Exit(0)
}
