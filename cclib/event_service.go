package cclib

import (
	"log"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	kb "github.com/philipjkim/kafka-brokers-go"
	"github.com/wvanbergen/kafka/consumergroup"
)

type EventHandler func(payload []byte)

type EventService struct {
	serviceID string
	zkNodes   []string

	kafkaProducer sarama.SyncProducer

	handlers map[string]EventHandler
}

func NewEventService(zkNodes []string, serviceID string) (*EventService, error) {
	svc := &EventService{
		serviceID: serviceID,
		zkNodes:   zkNodes,
		handlers:  make(map[string]EventHandler),
	}
	err := svc.setupKafkaProducer()
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func (svc *EventService) Register(event string, handler EventHandler) {
	svc.handlers[event] = handler
}

func (svc *EventService) Publish(event string, payload []byte) error {
	message := &sarama.ProducerMessage{Topic: event}
	message.Value = sarama.ByteEncoder(payload)
	_, _, err := svc.kafkaProducer.SendMessage(message)
	if err != nil {
		return err
	}
	log.Printf("Published event: %s\n", event)
	return nil
}

func (svc *EventService) setupKafkaProducer() error {
	kbConn, err := kb.NewConn(svc.zkNodes)
	if err != nil {
		return err
	}
	brokerList, _, err := kbConn.GetW()
	if err != nil {
		return err
	}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewManualPartitioner

	svc.kafkaProducer, err = sarama.NewSyncProducer(brokerList, config)
	return err
}

func (svc *EventService) Start() error {
	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetNewest
	config.Offsets.ProcessingTimeout = 10 * time.Second

	topics := svc.topics()
	log.Printf("Subscribing topics: %s\n", strings.Join(topics, ","))
	consumer, err := consumergroup.JoinConsumerGroup(
		svc.serviceID, topics, svc.zkNodes, config,
	)
	if err != nil {
		return err
	}
	go svc.listenKafkaConsumer(consumer)
	return nil
}

func (svc *EventService) listenKafkaConsumer(consumer *consumergroup.ConsumerGroup) {
	for message := range consumer.Messages() {
		if handler, ok := svc.handlers[message.Topic]; ok {
			log.Printf("Received event: %s\n", message.Topic)
			go handler(message.Value)
			consumer.CommitUpto(message)
		}
	}
}

func (svc *EventService) topics() []string {
	result := make([]string, 0, len(svc.handlers))
	for event := range svc.handlers {
		result = append(result, event)
	}
	return result
}
