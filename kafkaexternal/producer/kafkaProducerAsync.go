package producer

/*
 kafka 消息发送异步版本
*/
import (
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/bingtianyiyan/youki_gotools/kafkaexternal/kafkaconfig"
	"strings"
)

type ProducerAsync struct {
	producer sarama.AsyncProducer
}

func NewProducerAsync(config *sarama.Config,kafkaConfig kafkaconfig.KafkaConfig) *ProducerAsync {
	producer,err := InitKafkaProducerAsyncClient(config,kafkaConfig)
	if err != nil{
		return nil
	}

	// [!important] 异步生产者发送后,设置为true必须把返回值从 Errors 或者 Successes 中读出来 不然会阻塞 sarama 内部处理逻辑 导致只能发出去一条消息
	go func() {
		for s := range producer.Successes() {
			fmt.Println(s.Value)
		}
	}()

	go func() {
		for e := range producer.Errors() {
			if e != nil {
				fmt.Printf("[Producer] err:%v msg:%+v \n", e.Msg, e.Err)
			}
		}
	}()

	return &ProducerAsync{
		producer:  producer,
	}
}

// InitKafkaProducerAsyncClient 初始化客户端
func InitKafkaProducerAsyncClient(config *sarama.Config,kafkaConfig kafkaconfig.KafkaConfig) (sarama.AsyncProducer,error) {
	if kafkaConfig.Address == ""{
		panic("Producer Address is empty")
	}
	kafkaAddress := strings.Split(kafkaConfig.Address,",")
	//配置
	GetKafkaProducerConfig(config)
	// 连接kafka
	client, err := sarama.NewClient(kafkaAddress, config)
	if err != nil {
		fmt.Print(err)
	}
	producer, err := sarama.NewAsyncProducerFromClient(client)
	return producer,err
}

func (m *ProducerAsync) ProducerSendMsgAsync(topic string, value interface{}) error {
	var k_value sarama.Encoder
	switch value.(type) {
	case string:
		k_value = sarama.StringEncoder(value.(string))
	case int:
		k_value = sarama.ByteEncoder{byte(value.(int))}
	case byte:
		k_value = sarama.ByteEncoder{value.(byte)}
	default:
		return errors.New("msg error")
	}
	message := &sarama.ProducerMessage{Topic: topic, Value: k_value}
	m.producer.Input() <- message
	return nil
}


func (m *ProducerAsync) Close() {
	m.producer.AsyncClose()
}
