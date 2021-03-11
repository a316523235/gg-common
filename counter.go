package gg_common

type Counter interface {
	Incr(...interface{})
	GetKey(...interface{}) string
	RunConter()
}

type FilterCounter struct {
	key string
}

func (c FilterCounter) Incr(dspKey string, isOk int) {

}

func (c FilterCounter) GetKey(dspKey string, isOk int) string {
	return "xx"
}

func (c FilterCounter) RunConter() {

}
