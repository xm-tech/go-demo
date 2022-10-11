package main

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	_proto "github.com/xm-tech/go-demo/mq/kafka/proto"
)

// TODO 这里是 发 protobuf 结构体消息到 kafka 的例子， 我们的实际案例中发的是 JSONArray
func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	careerData0 := &_proto.CareerData{Idx: 50554300, Time: 1111111, Count: 1, ConditionId: 93} // protobuf Message Struct
	careerDataBytes, err := proto.Marshal(careerData0)
	if err != nil {
		fmt.Printf("Marshal Fail = %+v\n")
		return
	}
	msg := &sarama.ProducerMessage{}
	msg.Topic = "CAREER_DATA_TOPIC_3"
	msg.Value = sarama.ByteEncoder(careerDataBytes)

	// 同步连接 kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Printf("NewSyncProducer Err = %+v\n", err)
		return
	}
	defer client.Close()

	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Printf("SendMessage Fail,err = %+v\n", err)
		return
	}

	fmt.Printf("SendMesssage Succ ,pid=%+v,offset=%+v\n", pid, offset)

}
