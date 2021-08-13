package rabbitmq

const (
	ms  = 1
	s   = 1000
	min = 60000
)

type header map[string]interface{}

type body interface{}

type marshalFunc func(interface{}) ([]byte, error)

type handlerFunc func([]byte) error

type Message interface {
	MessageHeaderInit()
	MessageHeader() header
	MessageBody() body
	MessageRoutingKey() string
	MessagePriority() int
}

func delay(m Message, d int) {
	setHeader(m, header{
		XDelayHeader: d,
	})
}

func RemoveDelay(m Message) {
	delHeader(m, XDelayHeader)
}

func DelayMilisecond(m Message, d int) {
	delay(m, d*ms)
}

func DelaySecond(m Message, d int) {
	delay(m, d*s)
}

func DelayMinute(m Message, d int) {
	delay(m, d*min)
}
