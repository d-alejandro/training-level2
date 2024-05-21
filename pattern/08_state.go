package main

import "fmt"

/*
 Реализовать паттерн «состояние».
 Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования на практике.
 https://en.wikipedia.org/wiki/State_pattern

 Поведенческий паттерн

 Применимость:
 Инкапсулирует различные варианты поведения объекта в зависимости от его внутреннего состояния. Предоставляет
 возможность заменить использование условных операторов, в которых выбор ветви зависит от состояния.

 Плюсы:
 Упрощает поддержку кода.
 Упрощает добавление новых вариантов поведения объекта.
 Делает явными переходы между состояниями.

 Минусы:
 Усложняет код при наличии малого количества состояний.

 Реальное использование:
 При работе с протоколом установления TCP-соединений.
*/

type StateStatus int

const (
	StateAStatus StateStatus = iota
	StateBStatus
)

type State interface {
	Handle()
}

//------------------------------------------------------
//-- StateA --------------------------------------------

type StateA struct {
	stateContext *StateContext
}

func NewStateA(stateContext *StateContext) State {
	return &StateA{stateContext}
}

func (receiver *StateA) Handle() {
	fmt.Println("StateA:Handle")

	receiver.stateContext.SetState(StateBStatus)
}

//------------------------------------------------------
//-- StateB --------------------------------------------

type StateB struct {
	stateContext *StateContext
}

func NewStateB(stateContext *StateContext) State {
	return &StateB{stateContext}
}

func (receiver *StateB) Handle() {
	fmt.Println("StateB:Handle")

	receiver.stateContext.SetState(StateAStatus)
}

//------------------------------------------------------
//-- StateContext --------------------------------------

type StateContext struct {
	state State
}

func NewStateContext() *StateContext {
	stateContext := new(StateContext)
	stateContext.SetState(StateAStatus)
	return stateContext
}

func (receiver *StateContext) Request() {
	receiver.state.Handle()
}

func (receiver *StateContext) SetState(stateStatus StateStatus) {
	switch stateStatus {
	case StateAStatus:
		receiver.state = NewStateA(receiver)
	case StateBStatus:
		receiver.state = NewStateB(receiver)
	default:
		panic("unhandled default case")
	}
}

//------------------------------------------------------

/*
 * Client
 */
func main() {
	context := NewStateContext()
	context.Request()
	context.Request()
	context.Request()
}

/*
StateA:Handle
StateB:Handle
StateA:Handle
*/
