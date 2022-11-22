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
	"time"

	"github.com/arangodb/go-driver"
	amqp "github.com/rabbitmq/amqp091-go"
	pb "github.com/slntopp/nocloud-proto/services_providers"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

var (
	RabbitMQConn string
)

type PublicDataPubSub struct {
	log  *zap.Logger
	db   *driver.Database
	rbmq *amqp.Connection
}

func NewPublicDataPubSub(log *zap.Logger, db *driver.Database, rbmq *amqp.Connection) *PublicDataPubSub {
	ps := &PublicDataPubSub{
		log: log.Named("PublicDataServer"), rbmq: rbmq,
	}
	if db != nil {
		ps.db = db
	}
	return ps
}

func (s *PublicDataPubSub) Channel() *amqp.Channel {
	log := s.log.Named("Channel")

	ch, err := s.rbmq.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel", zap.Error(err))
	}
	return ch
}

func (s *PublicDataPubSub) TopicExchange(ch *amqp.Channel, name string) {
	err := ch.ExchangeDeclare(
		name, "topic", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatal("Failed to declare an exchange", zap.Error(err))
	}
}

func (s *PublicDataPubSub) PublicDataConsumerInit(ch *amqp.Channel, exchange, subtopic, col string) {
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

	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal("Failed to register a consumer", zap.Error(err))
	}
	go s.Consumer(col, msgs)
}

const updatePublicDataQuery = `
UPDATE DOCUMENT(@@collection, @key) WITH { public_data: @public_data } IN @@collection OPTIONS { mergeObjects: false }
`

func (s *PublicDataPubSub) Consumer(col string, msgs <-chan amqp.Delivery) {
	log := s.log.Named(col)
	log.Debug("PublicData updating started")
	for msg := range msgs {
		log.Debug("pd upd msg", zap.Any("msg", msg))
		var req pb.ObjectPublicData
		err := proto.Unmarshal(msg.Body, &req)
		if err != nil {
			log.Error("Failed to unmarshal request", zap.Error(err))
			if err = msg.Ack(false); err != nil {
				log.Warn("Failed to Acknowledge the delivery while unmarshal message", zap.Error(err))
			}
			continue
		}
		log.Debug("req pd", zap.Any("req", &req))
		c, err := (*s.db).Query(context.TODO(), updatePublicDataQuery, map[string]interface{}{
			"@collection": col,
			"key":         req.Uuid,
			"public_data": req.Data,
		})
		if err != nil {
			log.Error("Failed to update public_data", zap.Error(err))
			if err = msg.Nack(false, false); err != nil {
				log.Warn("Failed to Acknowledge the delivery while Update db", zap.Error(err))
			}
			continue
		}
		log.Debug("Updated public_data", zap.String("type", col), zap.String("uuid", req.Uuid))
		if err = c.Close(); err != nil {
			log.Warn("Error closing Driver cursor", zap.Error(err))
		}
		if err = msg.Ack(false); err != nil {
			log.Warn("Failed to Acknowledge the delivery", zap.Error(err))
		}
	}
}

type Pub func(msg *pb.ObjectPublicData) error

func (s *PublicDataPubSub) Publisher(ch *amqp.Channel, exchange, subtopic string) Pub {
	topic := exchange + "." + subtopic
	return func(msg *pb.ObjectPublicData) error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		body, err := proto.Marshal(msg)
		if err != nil {
			return err
		}
		return ch.PublishWithContext(ctx, exchange, topic, false, false, amqp.Publishing{
			ContentType:  "text/plain",
			DeliveryMode: amqp.Persistent,
			Body:         body,
		})
	}
}
