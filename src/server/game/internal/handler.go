package internal

import(
	"leaf/log"
	"leaf/gate"
	"server/msg"
	"reflect"
)

func handler(m interface{}, h interface{}){
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init(){
	handler(&msg.Hello{}, handleHello)
}

func handleHello(args[]interface{}){
	m := args[0].(*msg.Hello)
	a := args[1].(gate.Agent)

	log.Debug("hello %v", m.Name)

	a.WriteMsg(&msg.Hello{Name:"This is From Server"})
}
