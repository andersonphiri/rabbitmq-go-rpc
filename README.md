# rabbitmq-go-rpc
rpc design with rabbit mq and golang

## Use cases:
suppose you have a long running operation in which the result or response cannot be waited for within an http request lifetime window.
one approach is to send your request to the backend processor or service through a message bus and set the following parameters:
- routing_key / queuename
- correlationId
- replyTo queue name  - same value as the routing_key parameter

the backend processorwill then send results to the specified route and the client / consumer will filter response based on the correlation id and routing key specified in the original request. In this example we are using Factorial of a large number as our simulation for a long runing process
