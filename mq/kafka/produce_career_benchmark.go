package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/Shopify/sarama"
	_proto "github.com/xm-tech/go-demo/mq/kafka/proto"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	data, err := ioutil.ReadFile("kafka.conf")
	if err != nil {
		panic(err)
	}
	kafkaAddr := string(data)
	fmt.Println("kafka.addr=", kafkaAddr)

	client, err := sarama.NewSyncProducer([]string{kafkaAddr}, config)
	if err != nil {
		fmt.Printf("NewSyncProducer Err = %+v\n", err)
		return
	}
	defer client.Close()

	minIdx := int32(7000001)
	maxIdx := int32(8999999)
	// maxIdx := int32(7000002)

	now := time.Now().Unix()
	for id := minIdx; id < maxIdx; id++ {
		go sendMsg2Kafka(id, now, client)
	}
}

func sendMsg2Kafka(idx int32, time int64, kafkaClient sarama.SyncProducer) {
	careerData := &_proto.CareerData{Idx: idx, Time: time, Count: 1, ConditionId: 93} // protobuf Message Struct
	var careerDataArr []*_proto.CareerData
	careerDataArr = append(careerDataArr, careerData)
	arr, err := json.Marshal(careerDataArr)
	if err != nil {
		log.Printf("json.Marshal Err, input=%+v\n", careerDataArr)
		return
	}
	msg := &sarama.ProducerMessage{Topic: "CAREER_DATA_TOPIC_7", Value: sarama.StringEncoder(string(arr))}
	// 发送消息
	pid, offset, err := kafkaClient.SendMessage(msg)
	if err != nil {
		fmt.Printf("SendMessage Fail,err = %+v\n", err)
		return
	}
	log.Printf("SendMesssage Succ ,pid=%+v,offset=%+v,arr=%+v \n", pid, offset, string(arr))
}
