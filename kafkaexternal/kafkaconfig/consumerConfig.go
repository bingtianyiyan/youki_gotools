package kafkaconfig

import (
	"github.com/Shopify/sarama"
)

/*
消费者配置
 */

type ConsumerConfig struct {
	Topic string
	OffsetMode int64 //从最早消费还是最迟
	GroupName string
	AutoCommit bool
	RetErrors bool
	RetSuccess bool
	RetNotification bool
	ApiHandler sarama.ConsumerGroupHandler
}
