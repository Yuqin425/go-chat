package chat

import (
	"IM_QQ/settings"
	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"
)

var producer *nsq.Producer
var topic string = "default_message"

func InitProducer(cfg *settings.MsgChannelType) (err error) {
	topic = cfg.NsqTopic
	config := nsq.NewConfig()

	producer, err = nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		zap.L().Error("init nsq producer error", zap.Any("init nsq producer error", err.Error()))
	}
	return nil
}

func Send(data []byte) {
	producer.Publish(topic, data)
}

func Close() {
	if producer != nil {
		producer.Stop()
	}
}
