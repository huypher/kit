package rabbitmq

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func Test_channel_Consume(t *testing.T) {
	r1 := NewRabbitmq(Address(addr), Logging(true))
	c := r1.CreatePublisher(
		ExchangeName("exchange_direct"),
		ExchangeKind(ExchangeDirect),
		RegisterMarshalFunc(json.Marshal),
	)

	r2 := NewRabbitmq(Address(addr), Logging(true))
	q := r2.CreateConsumer(
		QueueName("queue_test1"),
		PriorityQueue(255),
		RegisterHandlerFunc(func(bytes []byte) error {
			fmt.Println(string(bytes))
			//time.Sleep(1 * time.Second)
			return nil
		}),
	)

	q.Bind("exchange_direct", "routing_key")

	messages := []Msg{}
	for i := 1; i <= 100; i++ {
		messages = append(messages, Msg{
			headers:    header{},
			Body:       i,
			RoutingKey: "routing_key",
			Priority:   i,
		})
	}

	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, msg := range messages {
				c.Publish(msg)
			}
		})
	}
	time.Sleep(5 * time.Second)
	q.Consume()
	for {
		continue
	}
}

func Test_Delay_Message(t *testing.T) {
	r1 := NewRabbitmq(Address(addr), Logging(true))
	c := r1.CreatePublisher(
		ExchangeName("exchange_direct"),
		ExchangeKind(ExchangeDirect),
		ExchangeDelayedType(),
		RegisterMarshalFunc(json.Marshal),
	)

	r2 := NewRabbitmq(Address(addr))
	q := r2.CreateConsumer(
		QueueName("queue_test1"),
		RegisterHandlerFunc(func(bytes []byte) error {
			fmt.Println(string(bytes))
			return nil
		}),
	)

	q.Bind("exchange_direct", "routing_key")

	msg1 := Msg{
		Body:       "msg1",
		RoutingKey: "routing_key",
	}
	msg1.DelaySecond(10)

	msg2 := Msg{
		Body:       "msg2",
		RoutingKey: "routing_key",
	}

	messages := []Msg{msg1, msg2}

	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, msg := range messages {
				c.Publish(msg)
			}
		})
	}
	//time.Sleep(20 * time.Second)
	q.Consume()
	for {
		continue
	}
}
