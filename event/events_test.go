package event

import(
	"reflect"
	"testing"
)

func assertDeepEqualT(t *testing.T) func(any, any, string) {
	return func(actual any, expected any, msg string) {
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Failed %s: actual: %v, expected: %v", msg, actual, expected)
		}
	}
}

func Test_Listen(t *testing.T) {
	eq := assertDeepEqualT(t)

    checkData := "index data"
    var eventData any

    Listen("data.test", func(data any) {
        eventData = data
    })
    Dispatch("data.test", checkData)

    eq(eventData, checkData, "Listen")

    // ==========

    checkData2 := "index data 222"
    var eventData2 any

    ev := NewEvents()

    ev.Listen("data.test111111", func(data any) {
        eventData2 = data
    })
    ev.Dispatch("data.test111111", checkData2)

    eq(eventData2, checkData2, "Listen2")

    // ==========

    checkData3 := "index data many"
    var eventData3, eventData3_1, eventData3_2 any

    Listen("many.test1", func(data any) {
        eventData3 = data
    })
    Listen("many.test2", func(data any) {
        eventData3_1 = data
    })
    Listen("many.test3", func(data any) {
        eventData3_2 = data
    })
    Dispatch("many.*", checkData3)

    eq(eventData3, checkData3, "Listen eventData3")
    eq(eventData3_1, checkData3, "Listen eventData3_1")
    eq(eventData3_2, checkData3, "Listen eventData3_2")
}

var testEventRes map[string]any

func init() {
    testEventRes = make(map[string]any)
}

type TestEvent struct {}

func (this *TestEvent) OnTestEvent(data any) {
    testEventRes["TestEvent_OnTestEvent"] = data
}

func (this *TestEvent) OnTestEventName(data any, name string) {
    testEventRes["TestEvent_OnTestEventName"] = data
    testEventRes["TestEvent_OnTestEventNameName"] = name
}

type TestEventPrefix struct {}

func (this TestEventPrefix) EventPrefix() string {
    return "ABC"
}

func (this TestEventPrefix) OnTestEvent(data any) {
    testEventRes["TestEventPrefix_OnTestEvent"] = data
}

type TestEventSubscribe struct {}

func (this *TestEventSubscribe) Subscribe(e *Events) {
    e.Listen("TestEventSubscribe", this.OnTestEvent)
}

func (this *TestEventSubscribe) OnTestEvent(data any) {
    testEventRes["TestEventSubscribe_OnTestEvent"] = data
}

type TestEventStructData struct {
    Data string
}

func EventStructTest(data TestEventStructData, name any) {
    testEventRes["EventStructTest"] = data.Data
    testEventRes["EventStructTest_Name"] = name
}

type TestEventStructHandle struct {}

func (this *TestEventStructHandle) Handle(data any) {
    testEventRes["TestEventStructHandle_Handle"] = data
}

func Test_Subscribe(t *testing.T) {
	eq := assertDeepEqualT(t)

    checkData := "index data Test_Subscribe"

    Subscribe(&TestEvent{})
    Dispatch("TestEvent", checkData)
    Dispatch("TestEventName", checkData)

    eq(testEventRes["TestEvent_OnTestEvent"], checkData, "Subscribe 1")
    eq(testEventRes["TestEvent_OnTestEventName"], checkData, "Subscribe 2")
    eq(testEventRes["TestEvent_OnTestEventNameName"], "TestEventName", "Subscribe 2")

    // =======

    ev := NewEvents()

    checkData2 := "index data Test_Subscribe 2"

    ev.Subscribe(&TestEvent{})
    ev.Dispatch("TestEvent", checkData2)
    ev.Dispatch("TestEventName", checkData2)

    eq(testEventRes["TestEvent_OnTestEvent"], checkData2, "Subscribe 2-1")
    eq(testEventRes["TestEvent_OnTestEventName"], checkData2, "Subscribe 2-2")
    eq(testEventRes["TestEvent_OnTestEventNameName"], "TestEventName", "Subscribe Name 2-2")
}

func Test_Subscribe_Prefix(t *testing.T) {
	eq := assertDeepEqualT(t)

    checkData := "index data Test_Subscribe_Prefix"

    Subscribe(TestEventPrefix{})
    Dispatch("ABCTestEvent", checkData)

    eq(testEventRes["TestEventPrefix_OnTestEvent"], checkData, "Subscribe 1")

    // =======

    ev := NewEvents()

    checkData2 := "index data Test_Subscribe_Prefix 2"

    ev.Subscribe(TestEventPrefix{})
    ev.Dispatch("ABCTestEvent", checkData2)

    eq(testEventRes["TestEventPrefix_OnTestEvent"], checkData2, "Subscribe 2-1")
}

func Test_EventSubscribe(t *testing.T) {
	eq := assertDeepEqualT(t)

    checkData := "index data Test_EventSubscribe"

    Subscribe(&TestEventSubscribe{})
    Dispatch("TestEventSubscribe", checkData)

    eq(testEventRes["TestEventSubscribe_OnTestEvent"], checkData, "Subscribe 1")

    // =======

    ev := NewEvents()

    checkData2 := "index data Test_EventSubscribe 2"

    ev.Subscribe(&TestEventSubscribe{})
    ev.Dispatch("TestEventSubscribe", checkData2)

    eq(testEventRes["TestEventSubscribe_OnTestEvent"], checkData2, "Subscribe 2-1")
}

func Test_EventStruct(t *testing.T) {
	eq := assertDeepEqualT(t)

    checkData := "index data Test_EventStruct"

    Listen(TestEventStructData{}, EventStructTest)
    Dispatch(TestEventStructData{
        Data: checkData,
    })

    eq(testEventRes["EventStructTest"], checkData, "Subscribe 1")
    eq(testEventRes["EventStructTest_Name"], "github.com/deatil/go-event/event.TestEventStructData", "Subscribe Name 2-2")

    // =======

    ev := NewEvents()

    checkData2 := "index data Test_EventStruct 2"

    ev.Listen(TestEventStructData{}, EventStructTest)
    ev.Dispatch(TestEventStructData{
        Data: checkData2,
    })

    eq(testEventRes["EventStructTest"], checkData2, "Subscribe 2-1")
    eq(testEventRes["EventStructTest_Name"], "github.com/deatil/go-event/event.TestEventStructData", "Subscribe Name 2-2")
}

func Test_EventStructHandle(t *testing.T) {
	eq := assertDeepEqualT(t)

    checkData := "index data Test_EventStructHandle"

    Listen("TestEventStructHandle", &TestEventStructHandle{})
    Dispatch("TestEventStructHandle", checkData)

    eq(testEventRes["TestEventStructHandle_Handle"], checkData, "Subscribe 1")

    // =======

    ev := NewEvents()

    checkData2 := "index data Test_EventStructHandle 2"

    ev.Listen("TestEventStructHandle", &TestEventStructHandle{})
    ev.Dispatch("TestEventStructHandle", checkData2)

    eq(testEventRes["TestEventStructHandle_Handle"], checkData2, "Subscribe 2-1")
}


