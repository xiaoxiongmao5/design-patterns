package service_a

import(
	"fmt"
	"xj/publish_subscribe/event_bus"
)

func A(){
	fmt.Println("我是A")
	fmt.Println("A:发布消息：玩家升级了")
	event_bus.GlobalEventBus.Emit("user upgraded", "xj", 5)
}