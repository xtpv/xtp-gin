package events

import (
	"context"
	"log"
)

// EventHandler 事件处理函数
type EventHandler func(arg interface{}) error

// Event 事件
type Event struct {
	ctx      context.Context
	eventCh  chan *eventData
	eventMap map[string]EventHandler
}

// eventData 事件内容
// Name 事件名称 Param 事件处理函数参数
type eventData struct {
	Name  string
	Param interface{}
}

// New 创建事件管理器
func New(ctx context.Context) *Event {
	return &Event{
		ctx:      ctx,
		eventCh:  make(chan *eventData),
		eventMap: make(map[string]EventHandler),
	}
}

// Bind 绑定事件
// name 事件名称 handler 事件处理函数
func (e *Event) Bind(name string, handler EventHandler) {
	e.eventMap[name] = handler
}

// Trigger 触发事件
// name 事件名称 param 事件处理函数的参数
func (e *Event) Trigger(name string, param interface{}) {
	e.eventCh <- &eventData{
		Name:  name,
		Param: param,
	}
}

// Start 启动事件触发器
func (e *Event) Start() {
	go func() {
		for {
			select {
			case <-e.ctx.Done():
				log.Println("event manager done!")
				return
			case eventData, ok := <-e.eventCh:
				if !ok {
					log.Println("event channel unexpectedly closed")
					return
				}
				h, ok := e.eventMap[eventData.Name]
				if !ok {
					log.Printf("can`t find event: %s", eventData.Name)
					return
				}
				// 使用协程处理
				go func() {
					err := h(eventData.Param)
					if err != nil {
						log.Println(err)
					}
				}()
			}
		}
	}()
}
