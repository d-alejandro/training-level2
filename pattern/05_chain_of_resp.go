package main

import "fmt"

/*
 Реализовать паттерн «цепочка вызовов».
 Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования на практике.
 https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

 Поведенческий паттерн

 Применимость:
 Связывает объекты получатели в цепочку и передает запрос вдоль этой цепочки, пока его не обработают.

 Плюсы:
 Динамическое формирование цепочки клиентом.
 Уменьшает степень связанности между клиентом и обработчиками.

 Минусы:
 Усложняет процесс мониторинга и отладки.

 Реальное использование:
 В API Cocoa и фреймворке Cocoa Touch, которые используются для приложений OS X и iOS соответственно.
*/

type Handler interface {
	SetNextHandler(Handler)
	Execute()
}

//-- HandlerA ------------------------------------------

type HandlerA struct {
	nextHandler Handler
}

func NewHandlerA() Handler {
	return &HandlerA{}
}

func (receiver *HandlerA) SetNextHandler(nextHandler Handler) {
	receiver.nextHandler = nextHandler
}

func (receiver *HandlerA) Execute() {
	fmt.Println("HandlerA:Execute")

	if receiver.nextHandler != nil {
		receiver.nextHandler.Execute()
	}
}

//------------------------------------------------------
//-- HandlerB ------------------------------------------

type HandlerB struct {
	nextHandler Handler
}

func NewHandlerB() Handler {
	return &HandlerB{}
}

func (receiver *HandlerB) SetNextHandler(nextHandler Handler) {
	receiver.nextHandler = nextHandler
}

func (receiver *HandlerB) Execute() {
	fmt.Println("HandlerB:Execute")

	if receiver.nextHandler != nil {
		receiver.nextHandler.Execute()
	}
}

//------------------------------------------------------
//-- HandlerC ------------------------------------------

type HandlerC struct {
	nextHandler Handler
}

func NewHandlerC() Handler {
	return &HandlerC{}
}

func (receiver *HandlerC) SetNextHandler(nextHandler Handler) {
	receiver.nextHandler = nextHandler
}

func (receiver *HandlerC) Execute() {
	fmt.Println("HandlerC:Execute")

	if receiver.nextHandler != nil {
		receiver.nextHandler.Execute()
	}
}

//------------------------------------------------------

/*
 * Client
 */
func main() {
	handlerA := NewHandlerA()
	handlerB := NewHandlerB()
	handlerC := NewHandlerC()

	handlerA.SetNextHandler(handlerB)
	handlerB.SetNextHandler(handlerC)

	handlerA.Execute()
}

/*
HandlerA:Execute
HandlerB:Execute
HandlerC:Execute
*/
