package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

var topic = "message-topic"
var brokerAddress = "my-kafka.default.svc.cluster.local:9092" //"192.168.31.204:9092" //

func SendMessage() error {

	// 创建 Kafka writer 实例
	/*writer1 := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{brokerAddress},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})*/
	writer := &kafka.Writer{
		Addr:         kafka.TCP(brokerAddress),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
		Async:        true,
	}
	defer writer.Close()
	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("Key-%d", time.Now().Unix())),
		Value: []byte("This is a test message from Module A."),
	}

	// 发送消息
	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		fmt.Println("Failed to send message:", err.Error())
		return err
	}
	fmt.Println("Sent message to topic:", topic)
	return nil
}
func ReadMessage() {
	// 定义 Kafka 服务器地址和 topic
	// 创建 Kafka reader 实例
	fmt.Println("now readmessage")
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{brokerAddress},
		Topic:     topic,
		Partition: 0,
		GroupID:   "module-b-group",
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	defer reader.Close()
	//for {
	// 使用ctx控制下超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	msg, err := reader.ReadMessage(ctx)
	if err != nil {
		fmt.Println("Failed to read message:", err.Error())
		return
	}
	fmt.Printf("Received message: %s = %s\n", string(msg.Key), string(msg.Value))
	return
	//	}
}
