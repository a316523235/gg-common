package notify

type SendType int

const SMS SendType = 1
const EMAIL SendType  = 2
const DINGTALK SendType  = 3

const MaxQueue  = 200 


type SendTo struct {
	SendType SendType
	SendAddr string
}

type Message struct {
	Title     string
	SubTitle  string
	Content   string
	DingDing  string
	EmailAddr string
	TellPhone string
}

type Notify interface {
	Send(msg Message)
}

