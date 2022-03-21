package producer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/bingtianyiyan/youki_gotools/kafkaexternal/kafkaconfig"
	"testing"
	"time"
)

func TestProducerSendMsg(t *testing.T){
	//全局实例化一次Producer即可，持有句柄
	var config *sarama.Config
	var producer =NewProducer(config,kafkaconfig.KafkaConfig{
		Address: "127.0.0.1:9092",
	})
	producer.ProducerSendMsg("MyTest1","HelloProducer")
}

func TestProducerAsync_ProducerSendMsgAsync(t *testing.T) {
	var config *sarama.Config
	var producer = NewProducerAsync(config,kafkaconfig.KafkaConfig{
		Address: "127.0.0.1:9092",
	})
	for i:=0;i<10;i++ {
		j := i
		go producer.ProducerSendMsgAsync("MyTest1",fmt.Sprintf("HelloProducerAsync->%v",j))
	}
	time.Sleep(time.Second *10)
}
