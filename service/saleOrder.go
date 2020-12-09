package service

import (
    "fmt"
    "github.com/Shopify/sarama"
    "sync"
    "gin-example/model"
    "gin-example/mq"
    "time"
)

func init() {

}

type SaleOrder struct {
    Id           int
    InvCode      string
    InvId        int
    InvName      string
    InvType      int
    Number       int
    UnitPrice    float32
    TotalPrice   float32
    CustomerId   int
    CustomerName string
    Remark       string
    Type         int
    OperatorId   int
    OperatorName string
}

func (this *SaleOrder) List() ([]*model.SaleOrder, error) {
    var saleOrderM model.SaleOrder
    list, err := saleOrderM.List()
    if err != nil {
        return nil, err
    }
    return list, nil
}
func (this *SaleOrder) Create() (int, error) {
    t1 := time.Now()
    var wg sync.WaitGroup
    wg.Add(2000)
    for i := 0; i < 2000; i++ {
        go func(this *SaleOrder) {
            defer wg.Done()
            saleOrder := model.SaleOrder{
                InvCode:      this.InvCode,
                InvId:        this.InvId,
                InvName:      this.InvName,
                InvType:      this.InvType,
                Number:       this.Number,
                UnitPrice:    this.UnitPrice,
                TotalPrice:   float32(this.Number) * this.UnitPrice,
                CustomerId:   this.CustomerId,
                CustomerName: this.CustomerName,
                Remark:       this.Remark,
                Type:         this.Type,
                OperatorId:   1,
                OperatorName: "默认操作人",
            }
            id, err := saleOrder.Create()
            fmt.Println(id, err)

        }(this)

    }
    wg.Wait()
    elapsed := time.Since(t1)
    fmt.Println("App elapsed: ", elapsed)

    /*if err != nil {
        return 0, err
    }*/
    return 0, nil

}
func (this *SaleOrder) Update() error {
    saleOrderM := model.SaleOrder{
        InvCode:      this.InvCode,
        InvId:        this.InvId,
        InvName:      this.InvName,
        InvType:      this.InvType,
        Number:       this.Number,
        UnitPrice:    this.UnitPrice,
        TotalPrice:   float32(this.Number) * this.UnitPrice,
        CustomerId:   this.CustomerId,
        CustomerName: this.CustomerName,
        Remark:       this.Remark,
        Type:         this.Type,
        OperatorId:   1,
        OperatorName: "默认操作人",
    }
    err := saleOrderM.Update(this.Id)
    if err != nil {
        return err
    }
    return nil

}

func (this *SaleOrder) Delete() error {
    var saleOrderM model.SaleOrder
    err := saleOrderM.Delete(this.Id)
    if err != nil {
        return err
    }
    return nil
}
func (this *SaleOrder) KafkaProducer() error {
    // 异步
    producer := mq.GetProducer()
    defer producer.Close()

    var value string
    for i := 0; i < 2; i++ {
        time.Sleep(500 * time.Millisecond)
        value = "this is a message 0606 _" + string(i)

        msg := &sarama.ProducerMessage{
            Topic: "test-topic-1",
            // Key:   sarama.StringEncoder("one_test"),
            Value: sarama.ByteEncoder(value),
        }
        // 发送消息
        pid, offset, err := producer.SendMessage(msg)
        if err != nil {
            fmt.Println("send msg failed, err:", err)
            return err
        }
        fmt.Printf("pid:%v offset:%v\n", pid, offset)
        // fmt.Println(value)
    }

    return nil
}

func (this *SaleOrder) KafkaAsyncProducer() error {
    // 异步
    producer := mq.GetAsyncProducer()
    go redChan(producer,456)
    go sendChan(producer)
    //

    return nil
}

func testD(p sarama.AsyncProducer) {
    fmt.Println("kafka关闭了")
    p.AsyncClose()
}
func sendChan(producer sarama.AsyncProducer) {
    var value string
    for i := 0; i < 456; i++ {
        //time.Sleep(500 * time.Millisecond)
        value = "this is a message 0606 _" + string(i)
        msg := &sarama.ProducerMessage{
            Topic: "test-topic-1",
            // Key:   sarama.StringEncoder("one_test"),
            Value: sarama.ByteEncoder(value),
        }
        // fmt.Println(value)
        // 使用通道发送
        producer.Input() <- msg
    }

}
func redChan(p sarama.AsyncProducer,count int) {
    defer testD(p)
    flag := 0;
    for {
        success := p.Successes()
        error := p.Errors()
        select {
        case <-success:
            flag++;
            fmt.Println("读到一个成功结果",flag)

        case fail := <-error:
            flag++;
            fmt.Println("读到一个成功结果", fail.Err,flag)

        default:
            //fmt.Println("这个不知道是啥")
        }
        if (flag>=count) {
            fmt.Println("协程管道监听完了")
            return
        }
    }

}
