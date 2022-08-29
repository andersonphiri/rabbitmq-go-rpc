package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/andersonphiri/rabbit-mq/remote-proc-calls/common"
	"github.com/andersonphiri/rabbit-mq/remote-proc-calls/models"
)
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
//var appSettings map[string]string

func main() {
	// create a connection
	appSettings, errLoad  := common.LoadMultipleAppSettingsFiles("rabbit.env")
	failOnError(errLoad, "no .env file found")
	mb := common.NewRabbitmqMessageBus(appSettings["rabbithost"],
	appSettings["rabbituser"], appSettings["rabbipass"], appSettings["rabbitport"], appSettings["rabbitprotocol"],
	)
	conn, err := mb.CreateTCPConnection()
	failOnError(err, "failed to create connection to amqp")
	defer conn.Close()

	// create channel 

	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()

	// set prefetc counts
	err = ch.Qos(
		1 , // prefetch count 
		0 , // prefetch size 
		false , // global
	)
	failOnError(err, "failed to set QOS")

	// declare a queue where requests will come through
	queue, err := ch.QueueDeclare(
		"rpc_queue", // name 
		false , // durable 
		false , // delete when unused 
		false , // exclusive 
		false , // no wait
		 nil , // 
	)

	failOnError(err, "failed to declare queue")

	mesages, err := ch.Consume(
		queue.Name,
		"" , // consumer 
		false, // auto-ack 
		false , //exclusive 
		false , // no-local 
		false , // no-wait 
		nil , // args 
	)
	failOnError(err, "failed to register consumer")

	var infinityLoop chan struct {}
	var wg sync.WaitGroup
	go func() {
		for message := range mesages {
			wg.Add(1)
			go func (message *amqp.Delivery)  {
				
				input, err := common.DeserializeTo[models.Request](message.Body)
				failOnError(err, fmt.Sprintf("failed to deserialize %s to Request type", message.Body))
				
				var respose = HandleRequest(input)
				resposeData, errSerial := json.Marshal(*respose)
				failOnError(errSerial, fmt.Sprintf("failed to deserialize response for request object %v", *input))
				// then publish message
				ctx := context.Background()
				err = ch.PublishWithContext(
					ctx, // context 
					"" , // exchange 
					message.ReplyTo, // routing key 
					false , //  mandatory 
					false , // immediate 
					amqp.Publishing{ ContentType: "application/json", CorrelationId: message.CorrelationId, Body: resposeData },
				)
				failOnError(err, fmt.Sprintf("failed to publish response for request object %v", *input))
				err = message.Ack(false)
				failOnError(err, "failed during acknowledge response")
				wg.Done()
				
			}(&message)

		}
		wg.Wait()

	}()

	log.Printf(" [*] Awaiting RPC requests")
	// https://www.rabbitmq.com/tutorials/tutorial-six-go.html

	<- infinityLoop


}

func toInt(param interface{}) (int, error)  {
	return strconv.Atoi(fmt.Sprintf("%v", param))
}

func printType(param interface{}) {
	switch t := param.(type) {
	case int:
		log.Printf("int %v\n", t)
	case int32:
		log.Printf("int32 %v\n", t)
	case int64:
		log.Printf("int32 %v\n", t)
	case int8:
		log.Printf("int32 %v\n", t)
	}

}

func HandleRequest(request *models.Request) *models.Response {
	var respose =  models.NewResponse()
	var meta = models.NewDefaultOperationMetadata()
	if request.Parameters == nil {
		respose.AddError(models.NewEmptyParametersError(common.FibonacciOperation,1))
		return respose
	}
	log.Println( "request parameter value is: ",request.Parameters)
	printType(request.Parameters)
	switch request.Operation {
	case common.FibonacciOperation:
		
		input, ok := toInt(request.Parameters)
		log.Printf("casted input value was: %d\n", input)
		if ok != nil {
			respose.AddError(models.NewCastParametersError(common.FibonacciOperation,"int"))
			respose.AddError(fmt.Errorf("%w",ok))
			break
		}

		res := FibInt64(input)
		respose.SetResult(res)
		meta.SetSpace("O(N)")
		meta.SetRuningTime("O(N)")
	case common.Factorial:
		input, ok := toInt(request.Parameters)
		log.Printf("casted input value was: %d\n", input)
		if ok != nil {
			respose.AddError(models.NewCastParametersError(common.FibonacciOperation,"int"))
			respose.AddError(fmt.Errorf("%w",ok))
			break
		}
		log.Printf("computing %s with input(s) %v\n", request.Operation,input)
		
		res := FactorialLarge(input)
		respose.SetResult(StringArray(res).String())
		meta.SetSpace("O(number of digits. see SterlingApproximationInt(N))")
		meta.SetRuningTime("O(number of digits. see SterlingApproximationInt(N))")
	default:
		respose.AddError(models.NewCastParametersError(common.FibonacciOperation,"int"))
		respose.AddError(fmt.Errorf("operation %s not yet implemented, sorry", request.Operation))	
	}
	
	respose.SetOperationsMetadata(meta)
	return respose
}




