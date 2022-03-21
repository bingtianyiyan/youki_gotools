package consumer

import (
	"context"
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/bingtianyiyan/youki_gotools/kafkaexternal/kafkaconfig"
	cluster "github.com/bsm/sarama-cluster"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"
)

//使用自定义方式消费,
func ConsumerHandlerCluster(consumerConfig kafkaconfig.ConsumerConfig,kafkaConfig kafkaconfig.KafkaConfig) error {
	if consumerConfig.Topic == ""{
		return errors.New("topic is empty!")
	}else if kafkaConfig.Address == ""  {
		return errors.New("address is empty!")
	}
	kafkaAddress := strings.Split(kafkaConfig.Address, ",")
	config := sarama.NewConfig()
	// Version 必须大于等于  V0_10_2_0
	config.Version = sarama.V0_10_2_1
	config.Consumer.Return.Errors = consumerConfig.RetErrors//返回错误则需要建立channnel监听
	config.Consumer.Offsets.Initial = consumerConfig.OffsetMode//消费从最早还是最迟
	config.Consumer.Offsets.AutoCommit.Enable = consumerConfig.AutoCommit//手动提交还是自动提交
	fmt.Println("start connect kafka")
	// 开始连接kafka服务器
	client, err := sarama.NewClient(kafkaAddress, config)
	if err != nil {
		return err
	}
	defer func() { _ = client.Close() }()

	group, err := sarama.NewConsumerGroupFromClient(consumerConfig.GroupName, client)
	if err != nil {
		return err
	}
	// 检查错误
	go func() {
		for err := range group.Errors() {
			fmt.Println("group errors : ", err)
		}
	}()

	ctx := context.Background()
	// for 是应对 consumer rebalance
	for {
		// 需要监听的主题
		topics := []string{consumerConfig.Topic}
		// 启动kafka消费组模式，消费的逻辑在上面的 ConsumeClaim 这个方法里
		err := group.Consume(ctx, topics, consumerConfig.ApiHandler)

		if err != nil {
			fmt.Println("consume failed; err : ", err)
			return err
		}
	}
}


func ConumerCluster(consumerConfig kafkaconfig.ConsumerConfig,kafkaConfig kafkaconfig.KafkaConfig) error{
	if consumerConfig.Topic == ""{
		return errors.New("topic is empty!")
	}else if kafkaConfig.Address == ""  {
		return errors.New("address is empty!")
	}
	kafkaAddress := strings.Split(kafkaConfig.Address, ",")

	config := cluster.NewConfig()
	config.Group.Return.Notifications = consumerConfig.RetNotification
	config.Consumer.Return.Errors = consumerConfig.RetErrors
	config.Consumer.Offsets.Initial = consumerConfig.OffsetMode
	config.Consumer.Offsets.AutoCommit.Enable = consumerConfig.AutoCommit
	config.Consumer.Offsets.CommitInterval = time.Second // 特么。。不加这个会报 non-positive interval for NewTicker
	//可以订阅多个主题
	topics := []string{consumerConfig.Topic}
	consumer, err := cluster.NewConsumer(kafkaAddress, consumerConfig.GroupName, topics, config)
	if err != nil {
		return err
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// 循环从通道中获取message
	//msg.Topic 消息主题
	//msg.Partition  消息分区
	//msg.Offset
	//msg.Key
	//msg.Value 消息值
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				fmt.Printf("%s receive message %s---Partition:%d, Offset:%d, Key:%s, Value:%s\n", consumerConfig.GroupName, msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				consumer.MarkOffset(msg, "") // 上报offset
			}
		case <-signals:
			return nil
		}
	}

}