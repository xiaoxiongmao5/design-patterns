package event_bus

import (
	"fmt"
	_ "errors"
)

/**
Events = {
	"name1"=>[
		1=>func,
		2=>func,
		3=>func
	]
}
*/

// 事件结构体
type Event struct {
	Name string
	Data interface{}
}

var GlobalEventBus *EventBus

type EventBus struct {
	EventId int
	EventMap map[string]map[int]func(args ...interface{})
}

func (this *EventBus) MakeOfEventMap(eventname string) {
	if this.EventMap == nil {
		fmt.Println("EventMap is nil, need make")
		this.EventMap = make(map[string]map[int]func(args ... interface{}))
	}
	if this.EventMap[eventname] == nil {
		fmt.Printf("EventMap[%s] is nil, need make\n", eventname)
		this.EventMap[eventname] = make(map[int]func(args ...interface{}))
	}
}

// 订阅
func (this *EventBus) On(eventname string, fn func(args ...interface{})) (int, error){
	this.EventId++
	this.MakeOfEventMap(eventname)
	this.EventMap[eventname][this.EventId] = fn
	return this.EventId, nil
}

func (this *EventBus) Off(eventname string, eventid int) (bool, error){
	this.MakeOfEventMap(eventname)
	delete(this.EventMap[eventname], eventid)
	return true, nil
}

func (this *EventBus) Emit(eventname string, args ... interface{}) {
	this.MakeOfEventMap(eventname)
	l := len(this.EventMap[eventname])
	fmt.Printf("eventname[%s]的订阅者有[%d]个\n", eventname, l )
	if l == 0 {
		fmt.Println("没有用户订阅该消息")
		return
	}
	for _, v := range this.EventMap[eventname] {
		v(args...)
	}
}