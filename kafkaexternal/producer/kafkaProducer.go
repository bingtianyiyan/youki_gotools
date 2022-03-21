package producer
/*
 kafka 消息发送同步版本
*/
import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/bingtianyiyan/youki_gotools/kafkaexternal/kafkaconfig"
	"strings"
)

type Producer struct {
	producer sarama.SyncProducer
}

// NewProducer 创建实例
func NewProducer(config *sarama.Config,kafkaConfig kafkaconfig.KafkaConfig) *Producer{
	var producer, err =InitKafkaProducerClient(config,kafkaConfig)
	if err != nil {
		fmt.Println(err)
	}
	return &Producer{
		producer: producer,
	}
}


// InitKafkaProducerClient 初始化客户端
func InitKafkaProducerClient(config *sarama.Config,kafkaConfig kafkaconfig.KafkaConfig) (sarama.SyncProducer,error) {
	if kafkaConfig.Address == ""{
		panic("Producer Address is empty")
	}
	kafkaAddress := strings.Split(kafkaConfig.Address,",")
	//配置
	GetKafkaProducerConfig(config)
	// 连接kafka
	client, err := sarama.NewSyncProducer(kafkaAddress, config)
	return client,err
}


// GetKafkaProducerConfig 获取生产者Config配置
func GetKafkaProducerConfig(config *sarama.Config) *sarama.Config{
	if config == nil {
		config = new(sarama.Config)
		config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
		config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
		config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
		config.Producer.Return.Errors = true
	}
	return config
}

// ProducerSendMsg 消息发送
func (m *Producer) ProducerSendMsg(topic string,msgData string) error {
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(msgData)
	// 发送消息
	pid, offset, err := m.producer.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return err
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
	return nil
}


