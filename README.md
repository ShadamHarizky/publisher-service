# Publisher-service - Technical Test PT.EDOT

## Persyaratan

Sebelum menjalankan service, pastikan point2 di bawah terpenuhi:

- [Go](https://golang.org/) (Golang) - Versi 1.18 atau lebih 
- [RabbitMQ](https://www.rabbitmq.com/)
- [Redis](https://redis.com/)

## Memulai

1. clone repositori:

    ```bash
    git clone https://github.com/ShadamHarizky/publisher-service
    cd publisher-service
    ```

2. install dependensi:

    ```bash
    - go mod init github.com/ShadamHarizky/publisher-service

    - go mod tidy
    ```

3. Buatkan .env sesuai dengan .env-example:

    - Rubah detail koneksi RabbitMQ dan Redis dalam file env.
    - Pastikan RabbitMQ diinstal dan berjalan.
    - Pastikan Redis diinstal dan berjalan.

4. Jalankan Project:

    ```bash
    go run .
    ```

## Konfigurasi

Update env pada file `.env` dan sesuaikan dengan konfigurasi pada komputer masing2:
- `PUBLISHER_TYPE`: Type publisher yang akan di gunakan `redis/rabbitmq`
- `RABBITMQ_URL`: URL server RabbitMQ (default: `amqp://guest:guest@localhost:5672/`)
- `REDIS_ADDRESS`: URL server Redis (default: `localhost:6379`)
- `REDIS_KEY`: Redis key pub/sub atau channel (default: `testing`)
- `RABBITMQ_ROUTING_KEY`: Routing Key RabbitMQ (default: `your-routing-key`)
