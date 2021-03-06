package sms

import (
	"hanya-go/pkg/config"
	"sync"
)

// Message 是短信的结构体
type Message struct {
	Template string
	Data     map[string]string

	Content string
}

// SMS 是我们发送短信的操作类
type SMS struct {
	Driver Driver
}

// once 单例模式
var once sync.Once

// internalSMS 内部使用的 SMS 对象
var internalSMS *SMS

var mode string

// NewSMS 单例模式获取
func NewSMS() *SMS {
	once.Do(func() {
		mode = config.GetString("sms.default")
		if mode == "aliyun" {
			internalSMS = &SMS{
				Driver: &Aliyun{},
			}
		} else if mode == "tencent" {
			internalSMS = &SMS{
				Driver: &Tencent{},
			}
		}

	})

	return internalSMS
}

func (sms *SMS) Send(phone string, message Message) bool {

	if mode == "aliyun" {
		return sms.Driver.Send(phone, message, config.GetStringMapString("sms.aliyun"))
	} else if mode == "tencent" {
		return sms.Driver.Send(phone, message, config.GetStringMapString("sms.tencent"))
	} else {
		return false
	}
}
