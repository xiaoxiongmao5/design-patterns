package service_c

import (
	"fmt"
	"xj/publish_subscribe/event_bus"
)

func Do(acct string, level int) {
	fmt.Printf("C:收到消息：玩家[%s]的等级升级到了[%d]\n", acct, level)
}

var EventIdMap map[string]int

func C(){
	fmt.Println("我是C")
	fmt.Println("C:取消订阅消息：玩家升级了")
	event_bus.GlobalEventBus.Off("user upgraded", EventIdMap["user upgraded"])
}

func Init() {
	EventIdMap = make(map[string]int)
	fmt.Println("C:订阅消息：玩家升级了")
	id, _ := event_bus.GlobalEventBus.On("user upgraded", func(args ...interface{}){
		acct := args[0].(string)
		level := args[1].(int)
		Do(acct, level)
	})
	EventIdMap["user upgraded"] = id
	fmt.Println("C:成功订阅消息,事件id为", id)
}