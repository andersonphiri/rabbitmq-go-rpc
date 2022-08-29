module gitlab.com/chitataunga/rabbit-mq/remote-proc-calls/producer

go 1.18

require (
	github.com/rabbitmq/amqp091-go v1.4.0
	github.com/andersonphiri/rabbit-mq/remote-proc-calls/common v0.0.0-dev-0001
	github.com/andersonphiri/rabbit-mq/remote-proc-calls/models v0.0.0-dev-0001
)

require github.com/joho/godotenv v1.4.0 // indirect

// replace gitlab.com/chitataunga/rabbit-mq/remote-proc-calls/models => ../models

// replace gitlab.com/chitataunga/rabbit-mq/remote-proc-calls/common => ../common
