package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	_proto "github.com/xm-tech/go-demo/mq/kafka/proto"
)

// TODO 这里是发送JSONArray 消息到 Kafka 的案例
func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	now := time.Now().Unix()
	fmt.Printf("now = %+v \n", now)

	careerData0 := &_proto.CareerData{Idx: 50554310, Time: now, Count: 1, ConditionId: 93} // protobuf Message Struct
	var careerDataArr []*_proto.CareerData
	careerDataArr = append(careerDataArr, careerData0)
	arr, _ := json.Marshal(careerDataArr)
	fmt.Printf("arr = %+v\n", string(arr))

	msg := &sarama.ProducerMessage{}
	msg.Topic = "CAREER_DATA_TOPIC_7"
	msg.Value = sarama.StringEncoder(string(arr))

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
