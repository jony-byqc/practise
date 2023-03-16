package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

var wg sync.WaitGroup

func main() {
	consumer, err := sarama.NewConsumer([]string{"192.168.60.181:9092"}, nil)
	if err != nil {
		fmt.Println("consumersync connect err:", err)
		return
	}
	defer consumer.Close()

	//获取 kafka 主题
	partitions, err := consumer.Partitions("asd")
	if err != nil {
		fmt.Println("get partitions failed, err:", err)
		return
	}

	for _, p := range partitions {
		//sarama.OffsetNewest：从当前的偏移量开始消费，sarama.OffsetOldest：从最老的偏移量开始消费
		partitionConsumer, err := consumer.ConsumePartition("asd", p, sarama.OffsetOldest)
		if err != nil {
			fmt.Println("partitionConsumer err:", err)
			continue
		}
		wg.Add(1)
		go func() {
			for m := range partitionConsumer.Messages() {
				fmt.Printf("key: %s, text: %s, offset: %d\n", string(m.Key), string(m.Value), m.Offset)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
