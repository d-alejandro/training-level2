package main

import "fmt"

/*
 Реализовать паттерн «посетитель».
 Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования на практике.
 https://en.wikipedia.org/wiki/Visitor_pattern

 Поведенческий паттерн

 Применимость:
 Когда нужно выполнить какую-то операцию над всеми элементами сложной структуры объектов (дерева объектов) без
 существенного изменения кодов этих объектов.

 Плюсы:
 Упрощает добавление новых операций.
 Объединяет родственные операции в одной структуре (Visitor).
 Visitor может запоминать в себе какое-то состояние по мере обхода контейнера.

 Минусы:
 Нарушает инкапсуляцию.
 Усложняет добавление новых структур ConcreteElement.
 Не оправдан, если иерархия элементов часто меняется.

 Реальное использование:
 Используется в Open Inventor (IRIS Inventor) - библиотеке для разработки приложений трёхмерной графики.
*/

type Element interface {
	Accept(Visitor)
	GetName() string
}

type Visitor interface {
	VisitElementA(*ElementA)
	VisitElementB(*ElementB)
}

//-- ElementA ------------------------------------------

type ElementA struct {
	name string
}

func NewElementA() Element {
	return &ElementA{"ElementA"}
}

func (receiver *ElementA) Accept(visitor Visitor) {
	visitor.VisitElementA(receiver)
}

func (receiver *ElementA) GetName() string {
	return receiver.name
}

//------------------------------------------------------
//-- ElementB ------------------------------------------

type ElementB struct {
	name string
}

func NewElementB() Element {
	return &ElementB{"ElementB"}
}

func (receiver *ElementB) Accept(visitor Visitor) {
	visitor.VisitElementB(receiver)
}

func (receiver *ElementB) GetName() string {
	return receiver.name
}

//------------------------------------------------------
//-- VisitorA ------------------------------------------

type VisitorA struct {
}

func NewVisitorA() *VisitorA {
	return &VisitorA{}
}

func (receiver *VisitorA) VisitElementA(element *ElementA) {
	fmt.Println("VisitorA", "visit", element.GetName())
}

func (receiver *VisitorA) VisitElementB(element *ElementB) {
	fmt.Println("VisitorA", "visit", element.GetName())
}

//------------------------------------------------------
//-- VisitorB ------------------------------------------

type VisitorB struct {
}

func NewVisitorB() *VisitorB {
	return &VisitorB{}
}

func (receiver *VisitorB) VisitElementA(element *ElementA) {
	fmt.Println("VisitorB", "visit", element.GetName())
}

func (receiver *VisitorB) VisitElementB(element *ElementB) {
	fmt.Println("VisitorB", "visit", element.GetName())
}

//------------------------------------------------------

/*
 * Client
 */
func main() {
	elementA := NewElementA()
	elementB := NewElementB()

	visitorA := NewVisitorA()
	elementA.Accept(visitorA)
	elementB.Accept(visitorA)

	visitorB := NewVisitorB()
	elementA.Accept(visitorB)
	elementB.Accept(visitorB)
}

/*
VisitorA visit ElementA
VisitorA visit ElementB
VisitorB visit ElementA
VisitorB visit ElementB
*/
