package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Shopify/sarama"
)

var (
	wg sync.WaitGroup
)

func main() {
	consumer, err := sarama.NewConsumer(strings.Split("127.0.0.1:9092", ","), nil)
	if err != nil {
		fmt.Errorf("failed to start consumer,err:%s", err)
		return
	}

	partions, err := consumer.Partitions("room")
	if err != nil {
		fmt.Printf("failed to get the partions: %+v\n", err)
		return
	}

	fmt.Printf("partions = %+v\n", partions)

	for partion := range partions {
		pc, err := consumer.ConsumePartition("room", int32(partion), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partion = %+v, %s\n", partion, err)
			return
		}

		defer pc.AsyncClose()
		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
			defer wg.Done()
			for {
				select {
				case msg := <-pc.Messages():
					fmt.Printf("msg: Partion: %+v, Offset: %+v, Key:%+v, Value:%+v\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				}
			}
			// for msg := range pc.Messages() {
			// 	fmt.Printf("msg = %+v\n", msg)
			// 	// fmt.Printf("Partion: %+v, Offset: %+v, Key:%+v, Value:%+v\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			// }
		}(pc)
	}
	wg.Wait()
	consumer.Close()
}
