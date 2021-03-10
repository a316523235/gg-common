package notify

import "testing"

func TestSms_Send(t *testing.T) {
	msg := Message{
		Title:"我是标题",
		Content:"我是内容",
		TellPhone:"18650313554",
	}
	sms := Sms{}
	sms.Send(msg)
}
