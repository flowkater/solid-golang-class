package consumers

import (
	"context"
	"event-data-pipeline/pkg/logger"
	"event-data-pipeline/pkg/rabbitmq"
	"event-data-pipeline/pkg/sources"
)

// compile type assertion check
var _ Consumer = new(RabbitMQConsumerClient)
var _ ConsumerFactory = NewRabbitMQConsumerClient

// ConsumerFactory 에 rabbitmq 컨슈머를 등록
func init() {
	Register("rabbitmq", NewRabbitMQConsumerClient)
}

type RabbitMQConsumerClient struct {
	rabbitmq.Consumer
	sources.Source
}

func NewRabbitMQConsumerClient(config jsonObj) Consumer {
	//TODO: 1주차 과제입니다.
	consumer := rabbitmq.NewRabbitMQConsumer(config)

	client := &RabbitMQConsumerClient{
		Consumer: consumer,
		Source:   sources.NewRabbitMQSource(consumer),
	}
	return client
}

// Init implements Consumer
func (rc *RabbitMQConsumerClient) Init() error {
	//TODO: 1주차 과제입니다.
	var err error

	err = rc.CreateConsumer()
	if err != nil {
		return err
	}

	err = rc.QueueDeclare()
	if err != nil {
		return err
	}

	err = rc.ExchangeDeclare()
	if err != nil {
		return err
	}

	return nil
}

// Consume implements Consumer
func (rc *RabbitMQConsumerClient) Consume(ctx context.Context) error {
	//TODO: 1주차 과제입니다.
	logger.Infof("Consume")

	err := rc.Read(ctx)
	if err != nil {
		return err
	}

	return nil
}
