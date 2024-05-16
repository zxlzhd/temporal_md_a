package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

var topic = "message-topic"
var brokerAddress = "my-kafka.default.svc.cluster.local:9092"

func SendMessage() error {

	// 创建 Kafka writer 实例
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{brokerAddress},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("Key-%d", time.Now().Unix())),
		Value: []byte("This is a test message from Module A."),
	}

	// 发送消息
	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		return err
	}
	fmt.Println("Sent message to topic:", topic)
	return nil
}
func ReadMessage() {
	// 定义 Kafka 服务器地址和 topic
	// 创建 Kafka reader 实例
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{brokerAddress},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	//for {
	msg, err := reader.ReadMessage(context.Background())
	if err != nil {
		fmt.Println("Failed to read message:", err)
		return
	}
	fmt.Printf("Received message: %s = %s\n", string(msg.Key), string(msg.Value))
	//	}
}
