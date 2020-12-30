package mq

import (
    "fmt"
    "github.com/Shopify/sarama"
    "log"
    "gin-example/config"
    "time"
)


func init() {
    fmt.Println("kafka")




}
//获取同步生产者
func GetProducer() sarama.SyncProducer{
    configS := sarama.NewConfig()
    configS.Producer.Return.Successes = true
    configS.Producer.Timeout = 5 * time.Second
    producer, err := sarama.NewSyncProducer(config.Kafka.Addrs, configS)
    if err != nil {
        fmt.Println("kafka 同步 链接失败",err)
        log.Printf("sarama.NewSyncProducer err, message=%s \n", err)
        return nil
    }
    log.Println("kafka 同步 链接成功")
    fmt.Println("kafka 同步 链接成功")
    return producer;
}

// 获取异步生产者
func GetAsyncProducer() sarama.AsyncProducer{
    configS := sarama.NewConfig()
    //等待服务器所有副本都保存成功后的响应
    configS.Producer.RequiredAcks = sarama.WaitForAll
    //随机向partition发送消息
    configS.Producer.Partitioner = sarama.NewRandomPartitioner
    //是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
    configS.Producer.Return.Successes = true
    configS.Producer.Return.Errors = true
    //设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
    //注意，版本设置不对的话，kafka会返回很奇怪的错误，并且无法成功发送消息
    //config.Version = sarama.V0_10_0_1

    fmt.Println("start make producer")
    //使用配置,新建一个异步生产者
    producer, err := sarama.NewAsyncProducer(config.Kafka.Addrs, configS)
    if err != nil {
        fmt.Println("kafka 异步 链接失败",err)
        log.Printf("sarama.NewAsyncProducer err, message=%s \n", err)
        return nil
    }
    log.Println("kafka 异步 链接成功")
    fmt.Println("kafka 异步 链接成功")
    return producer;
}
