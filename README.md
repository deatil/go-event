## go Event or Event Subscribe

<p align="center">
<a href="https://pkg.go.dev/github.com/deatil/go-event"><img src="https://pkg.go.dev/badge/deatil/go-event.svg" alt="Go Reference"></a>
<a href="https://codecov.io/gh/deatil/go-event" >
 <img src="https://codecov.io/gh/deatil/go-event/graph/badge.svg?token=SS2Z1IY0XL"/>
</a>
<img src="https://goreportcard.com/badge/github.com/deatil/go-event" />
</p>

### Desc

*  go Event or Event Subscribe pkg

[中文](README_CN.md) | English


### Download

~~~go
go get -u github.com/deatil/go-event
~~~


### Get Starting

~~~go
package main

import (
    "fmt"
    "github.com/deatil/go-event/event"
)

type TestEvent struct {}

func (this *TestEvent) OnTestEvent(data any) {
    fmt.Println("TestEvent: ")
    fmt.Println(data)
}

func (this *TestEvent) OnTestEventName(data any, name string) {
    fmt.Println("TestEventName: ")
    fmt.Println(data)
    fmt.Println(name)
}

type TestEventPrefix struct {}

func (this TestEventPrefix) EventPrefix() string {
    return "ABC"
}

func (this TestEventPrefix) OnTestEvent(data any) {
    fmt.Println("TestEventPrefix: ")
    fmt.Println(data)
}

type TestEventSubscribe struct {}

func (this *TestEventSubscribe) Subscribe(e *event.Events) {
    e.Listen("TestEventSubscribe", this.OnTestEvent)
}

func (this *TestEventSubscribe) OnTestEvent(data any) {
    fmt.Println("TestEventSubscribe: ")
    fmt.Println(data)
}

type TestEventStructData struct {
    Data string
}

func TestEventStruct(data TestEventStructData, name any) {
    fmt.Println("TestEventStruct: ")
    fmt.Println(data.Data)
    fmt.Println(name)
}

type TestEventStructHandle struct {}

func (this *TestEventStructHandle) Handle(data any) {
    fmt.Println("TestEventStructHandle: ")
    fmt.Println(data)
}

func main() {
    // Listen
    event.Listen("data.error", func(data any) {
        fmt.Println(data)
    })

    // Dispatch
    eventData := "index data"
    event.Dispatch("data.error", eventData)

    // call prefix `data.` all listener
    event.Dispatch("data.*", eventData)

    // ==================

    // Subscribe
    event.Subscribe(&TestEvent{})
    event.Subscribe(TestEventPrefix{})
    event.Subscribe(&TestEventSubscribe{})

    // Subscribe call
    event.Dispatch("TestEvent", eventData)
    event.Dispatch("TestEventName", eventData)
    event.Dispatch("ABCTestEvent", eventData)
    event.Dispatch("TestEventSubscribe", eventData)

    // ==================

    // Listen
    event.Listen(TestEventStructData{}, TestEventStruct)

    // Dispatch
    eventData2 := "index data"
    event.Dispatch(TestEventStructData{
        Data: eventData2,
    })

    // ==================

    // Listen
    event.Listen("TestEventStructHandle", &TestEventStructHandle{})

    // Dispatch
    event.Dispatch("TestEventStructHandle", eventData)
}

~~~


### LICENSE

*  The library LICENSE is `Apache2`, using the library need keep the LICENSE.


### Copyright

*  Copyright deatil(https://github.com/deatil).
