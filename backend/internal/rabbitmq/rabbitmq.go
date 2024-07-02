// Copyright 2024 Robert Cronin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

type RMQClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

var Client RMQClient

func Init() {
	// Get url from config
	host := viper.GetString("rabbitmq.host")
	port := viper.GetString("rabbitmq.port")
	user := viper.GetString("rabbitmq.user")
	password := viper.GetString("RABBITMQ_PASSWORD") // empty if on local, injected by k8s if on prod
	if viper.GetString("env") == "development" {
		password = "guest"
	}
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)

	// Connect to RabbitMQ
	conn, err := amqp.Dial(url)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to RabbitMQ: %v", err))
	}

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		panic(fmt.Sprintf("Failed to open a channel: %v", err))
	}

	Client.conn = conn
	Client.channel = ch
}

func (c *RMQClient) Close() error {
	if err := c.channel.Close(); err != nil {
		return err
	}
	return c.conn.Close()
}

func (c *RMQClient) PublishMessage(queueName string, body interface{}) error {
	q, err := c.channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	ctx := context.Background()

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	err = c.channel.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonBody,
		})

	if err != nil {
		return err
	}

	log.Printf(" [x] Sent %s\n", body)
	return nil
}

func (c *RMQClient) ConsumeMessages(queueName string, handler func([]byte) error) error {
	q, err := c.channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	msgs, err := c.channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			if err := handler(d.Body); err != nil {
				log.Printf("Error processing message: %v", err)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages on queue %s. To exit press CTRL+C", queueName)
	return nil
}
