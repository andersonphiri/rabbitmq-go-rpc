package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"math/rand"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/andersonphiri/rabbit-mq/remote-proc-calls/common"
	cm "github.com/andersonphiri/rabbit-mq/remote-proc-calls/common"
	"github.com/andersonphiri/rabbit-mq/remote-proc-calls/models"
)

const (
	MaxFactorial = 10_000
	MinFactorial = 2
)

var request *models.Request

func Init() {
	request = models.NewRequest(cm.Factorial, 10)
	log.Printf("executing init func...\n")
}

func GenerateRandomRequest() (data []byte) {
	n := randInt(MinFactorial, MaxFactorial)
	// if request == nil {
	// 	request = models.NewRequest(cm.Factorial, n)
	// } else {
	// 	request.Update(cm.Factorial,n)
	// }
	request = models.NewRequest(cm.Factorial, n)
	data, er := cm.Serialize(request)
	failOnError(er, "failed to serialize request object")
	return
}

func failOnErrorTemplate(err error, fmtTemplate string, params ...any) {
	if err != nil {
		log.Panicf("%s: %s", err, fmt.Sprintf(fmtTemplate,params...))
	}
}
func failOnError(err error, message string) {
	if err != nil {
		log.Panicf("%s : %s", err, message)
	}
}


func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
			bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func SendRequest(data []byte) (result *models.Response) {
	// create a connection
	appSettings, errLoad  := common.LoadMultipleAppSettingsFiles("rabbit.env")
	failOnError(errLoad, "no .env file found")
	mb := common.NewRabbitmqMessageBus(appSettings["rabbithost"],
	appSettings["rabbituser"], appSettings["rabbipass"], appSettings["rabbitport"], appSettings["rabbitprotocol"],
	)
	conn, err := mb.CreateTCPConnection()
	failOnError(err, "failed to create connection to message bus")
	// create channel 

	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	corrId := randomString(32)
	ctx := context.Background()
	err = ch.PublishWithContext(
		ctx,
		"",          // exchange
		"rpc_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
				ContentType:   "text/plain",
				CorrelationId: corrId,
				ReplyTo:       q.Name,
				Body:          data,
		})
    failOnError(err, "Failed to publish a message")
	for d := range msgs {
		if corrId == d.CorrelationId {
				result, err = cm.DeserializeTo[models.Response](d.Body)
				failOnErrorTemplate(err, "failed to deserialize response object with data %s", d.Body)
				break
		}
	}

return

}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	data := GenerateRandomRequest()

	log.Printf(" [x] Requesting object %v\n", *request)
	response := SendRequest(data)
	if response != nil {
		log.Printf(" [.] Got Result: %v\n", response.Result)
	} else {
		log.Printf(" [.] Got null response...\n")
	}
	
}
