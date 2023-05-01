package chat

import (
	"IM_QQ/settings"
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

var consumer *nsq.Consumer

type ConsumerCallback func(data []byte)

type ConsumerHandler struct {
	Title string
}

// HandleMessage 是需要实现的处理消息的方法
func (c *ConsumerHandler) HandleMessage(msg *nsq.Message) (err error) {
	ConsumerNsqMsg(msg.Body)
	return
}

func InitConsumer(cfg *settings.MsgChannelType) (err error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	consumer, err = nsq.NewConsumer(cfg.NsqTopic, cfg.ChannelType, config)
	if err != nil {
		fmt.Printf("create consumer failed, err:%v\n", err)
		return
	}
	consumer.AddHandler(&ConsumerHandler{
		Title: "IM",
	})

	// if err := c.ConnectToNSQD(address); err != nil { // 直接连NSQD
	if err = consumer.ConnectToNSQLookupd("127.0.0.1:4161"); err != nil { // 通过lookupd查询
		return err
	}
	return nil
}

func CloseConsumer() {
	if consumer != nil {
		consumer.Stop()
	}
}
