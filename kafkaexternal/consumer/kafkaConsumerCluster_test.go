package consumer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/bingtianyiyan/youki_gotools/kafkaexternal/kafkaconfig"
	"testing"
	"time"
)

type ConsumerGroupHandler struct{}

func (ConsumerGroupHandler) Setup(sess sarama.ConsumerGroupSession) error {
	return nil
}

func (ConsumerGroupHandler) Cleanup(sess sarama.ConsumerGroupSession) error {
	return nil
}

// 这个方法用来消费消息的
func (h ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// 获取消息
	for msg := range claim.Messages() {
		fmt.Printf("demo3 receive message %s---Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		// 手动确认消息 将消息标记为已使用
		sess.MarkMessage(msg, "")
	}
	return nil
}

func TestConsumerCluster(t *testing.T) {
	var consumerConfig = kafkaconfig.ConsumerConfig{
         Topic: "MyTest1",
         OffsetMode:sarama.OffsetNewest,
         GroupName: "MyGp",
         ApiHandler: ConsumerGroupHandler{},
         RetErrors: true,
         RetNotification: true,
	}
	var kafkaConfig  = kafkaconfig.KafkaConfig{
		 Address: "127.0.0.1:9092",
	}
	//go ConsumerHandlerCluster(consumerConfig,kafkaConfig)
	go ConumerCluster(consumerConfig,kafkaConfig)
	time.Sleep(time.Second*10000)
	fmt.Println("finish")
}
