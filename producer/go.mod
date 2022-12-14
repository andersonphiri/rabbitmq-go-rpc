module github.com/andersonphiri/rabbit-mq/remote-proc-calls/producer

go 1.18

require (
	github.com/andersonphiri/rabbit-mq/remote-proc-calls/common v0.0.0-00010101000000-000000000000
	github.com/andersonphiri/rabbit-mq/remote-proc-calls/models v0.0.0-00010101000000-000000000000
	github.com/rabbitmq/amqp091-go v1.4.0
)

require github.com/joho/godotenv v1.4.0 // indirect

replace github.com/andersonphiri/rabbit-mq/remote-proc-calls/models => ../models

replace github.com/andersonphiri/rabbit-mq/remote-proc-calls/common => ../common
