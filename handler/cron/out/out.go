package out

import (
	"fmt"
	"gin-template/server/cron"
)

type Out struct {
}

var instance *Out

func init() {
	// register output
	instance = &Out{}
	cron.RegisterCronHandler(instance)
}

func (o Out) Register() map[string]func() {
	return map[string]func(){
		"out": o.OutHandler,
	}
}

func (o Out) OutHandler() {
	fmt.Println("out")
}
