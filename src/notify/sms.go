package notify

import "fmt"

type Sms struct {}

func (sms Sms) Send(msg Message) {
	fmt.Printf("向（%s）发送短信：%s", msg.TellPhone, msg.Content)
}