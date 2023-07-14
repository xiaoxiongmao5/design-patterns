package main

import (
	"fmt"
	"xj/publish_subscribe/event_bus"
	"xj/publish_subscribe/service_a"
	"xj/publish_subscribe/service_b"
	"xj/publish_subscribe/service_c"
)



func init() {
	event_bus.GlobalEventBus = &event_bus.EventBus{}
	service_b.Init()
	service_c.Init()
	fmt.Println("app init ok~~")
}

func main() {
	service_a.A()
	fmt.Println()

	service_b.B()
	fmt.Println()

	service_a.A() 
	fmt.Println()
	
	service_c.C() 
	fmt.Println()

	fmt.Println()
	service_a.A() 
}