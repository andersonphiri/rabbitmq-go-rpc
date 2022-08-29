package common

import (
	"crypto/tls"
	"fmt"
	"reflect"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type IMessageBus interface {
	CreateTCPConnection() (*amqp.Connection ,error)
	CreateTCPSConnection(config *tls.Config) (*amqp.Connection ,error)
}

type Event struct {
	timestamp time.Time
}
func NewEvent() *Event {
	e := &Event{timestamp: time.Now()}
	return e
}

func (e *Event) GetTimestamp() *time.Time {
	return &e.timestamp
}

type Message struct {
	messageType string 
}
func NewMessage[T any](child *T) *Message {
	m := &Message{}
	m.messageType = reflect.ValueOf(child).Type().String()
	return m
}

func (m *Message) GetMessageType() string {
	return m.messageType
}

type IEventHandler interface {
	Handle(event *Event)
}

type Command struct {
	*Message
	timestamp time.Time
}
func NewCommand() *Command {
	c := &Command{timestamp: time.Now()}
	return c
}

func (cmd *Command) GetTimestamp() *time.Time {
	return &cmd.timestamp
}


type RabbitmqMessageBus struct {
	host string 
	user string
	password string 
	port string 
	protocol string
	connectionUrl string
}

func NewRabbitmqMessageBus(hostname string, username string, password string, port string, protocol string) *RabbitmqMessageBus {
	bus := &RabbitmqMessageBus{host: hostname, user: username, password: password, port: port, protocol: protocol}
	bus.connectionUrl = fmt.Sprintf("%s://%s:%s@%s:%s/", bus.protocol, bus.user, bus.password, bus.host, bus.port)
	return bus
}

func (mb *RabbitmqMessageBus) CreateTCPConnection() (*amqp.Connection ,error) {
	return amqp.Dial(mb.connectionUrl)
}

func (mb *RabbitmqMessageBus) CreateTCPSConnection(config *tls.Config) (*amqp.Connection ,error) {
	return amqp.DialTLS(mb.connectionUrl, config)
}