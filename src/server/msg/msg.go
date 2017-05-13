package msg

import (
	"leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Hello{})
}

type Hello struct{
	Name string
}
