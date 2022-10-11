package main

import (
	"encoding/json"
	"log"
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

	// async demo
	asyncProduer, err := sarama.NewAsyncProducer([]string{"127.0.0.1"}, config)
	if err != nil {
		log.Printf("err = %+v\n", err)
		panic(err)
	}

	minIdx := int32(7000001)
	// maxIdx := int32(8999999)
	maxIdx := int32(7000002)

	for id := minIdx; id < maxIdx; id++ {
		// 构造一个消息
		now := time.Now().Unix()
		careerData := &_proto.CareerData{Idx: id, Time: now, Count: 1, ConditionId: 93} // protobuf Message Struct
		var careerDataArr []*_proto.CareerData
		careerDataArr = append(careerDataArr, careerData)
		arr, _ := json.Marshal(careerDataArr)
		if err != nil {
			log.Printf("json.Marshal Err, input=%+v\n", careerDataArr)
			continue
		}
		log.Printf("now = %+v, arr = %+v \n", now, string(arr))
		msg := &sarama.ProducerMessage{Topic: "CAREER_DATA_TOPIC_7", Value: sarama.StringEncoder(string(arr))}
		// 发送消息
		asyncProduer.Input() <- msg
	}
}
