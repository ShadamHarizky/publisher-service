package messaging

// Publisher adalah interface untuk fungsi publish ke berbagai third party event driven (RabbitMQ, Redis, dll.)
type Publisher interface {
	Publish(key string, message string) error
	Close()
}
