/*
Copyright © 2021-2022 Nikita Ivanovski info@slnt-opp.xyz

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package states

import (
	"context"
	"log"

	"github.com/arangodb/go-driver"
	pb "github.com/slntopp/nocloud/pkg/states/proto"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

var (
	RabbitMQConn string
)

type StatesPubSub struct {
	log  *zap.Logger
	db   *driver.Database
	rbmq *amqp.Connection
}

func NewStatesPubSub(log *zap.Logger, db *driver.Database, rbmq *amqp.Connection) *StatesPubSub {
	ps := &StatesPubSub{
		log: log.Named("StatesServer"), rbmq: rbmq,
	}
	if db != nil {
		ps.db = db
	}
	return ps
}

func (s *StatesPubSub) Channel() *amqp.Channel {
	log := s.log.Named("Channel")

	ch, err := s.rbmq.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel", zap.Error(err))
	}
	return ch
}

func (s *StatesPubSub) TopicExchange(ch *amqp.Channel, name string) {
	err := ch.ExchangeDeclare(
		name, "topic", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatal("Failed to declare an exchange", zap.Error(err))
	}
}

func (s *StatesPubSub) StatesConsumerInit(ch *amqp.Channel, exchange, subtopic, col string) {
	if s.db == nil {
		log.Fatal("Failed to initialize states consumer, database is not set")
	}
	topic := exchange + "." + subtopic
	q, err := ch.QueueDeclare(
		topic, false, false, true, false, nil,
	)
	if err != nil {
		log.Fatal("Failed to declare a queue", zap.Error(err))
	}

	err = ch.QueueBind(q.Name, topic, exchange, false, nil)
	if err != nil {
		log.Fatal("Failed to bind a queue", zap.Error(err))
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Failed to register a consumer", zap.Error(err))
	}
	go s.Consumer(col, msgs)
}

const updateStateQuery = `
UPDATE DOCUMENT(@@collection, @key) WITH { state: @state } IN @@collection OPTIONS { mergeObjects: false }
`

func (s *StatesPubSub) Consumer(col string, msgs <-chan amqp.Delivery) {
	log := s.log.Named(col)
	for msg := range msgs {
		log.Debug("st upd msg", zap.Any("msg", msg))
		var req pb.ObjectState
		err := proto.Unmarshal(msg.Body, &req)
		if err != nil {
			log.Error("Failed to unmarshal request", zap.Error(err))
			continue
		}
		log.Debug("req st", zap.Any("req", &req))
		c, err := (*s.db).Query(context.TODO(), updateStateQuery, map[string]interface{}{
			"@collection": col,
			"key":         req.Uuid,
			"state":       req.State,
		})
		if err != nil {
			log.Error("Failed to update state", zap.Error(err))
			continue
		}
		log.Debug("Updated state", zap.String("type", col), zap.String("uuid", req.Uuid))
		c.Close()
	}
}

type Pub func(msg *pb.ObjectState) error

func (s *StatesPubSub) Publisher(ch *amqp.Channel, exchange, subtopic string) Pub {
	topic := exchange + "." + subtopic
	return func(msg *pb.ObjectState) error {
		body, err := proto.Marshal(msg)
		if err != nil {
			return err
		}
		return ch.Publish(exchange, topic, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	}
}
