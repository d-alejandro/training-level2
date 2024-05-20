package main

import "fmt"

/*
 Реализовать паттерн «фабричный метод».
 Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования на практике.
 https://en.wikipedia.org/wiki/Factory_method_pattern

 Порождающий паттерн

 Применимость:
 Применяется для создания объектов без указания их точной реализации. Делегирует создание объектов Factory.

 Плюсы:
 Уменьшает степень связанности между клиентом и продуктом.
 Реализует принцип открытости/закрытости.
 Упрощает поддержку кода.
 Упрощает добавление новых продуктов.

 Минусы:
 Необходимость создавать отдельную фабрику для каждого нового типа продукта.

 Реальное использование:
 В HTML5 DOM API интерфейс Document содержит фабричный метод createElement для создания определенных элементов
 интерфейса HTMLElement.
*/

type ProductFactory interface {
	Make() Product
}

type Product interface {
	Create()
}

//-- ProductA ------------------------------------------

type ProductA struct {
}

func NewProductA() Product {
	return &ProductA{}
}

func (receiver *ProductA) Create() {
	fmt.Println("ProductA:Create")
}

//------------------------------------------------------
//-- ProductB ------------------------------------------

type ProductB struct {
}

func NewProductB() Product {
	return &ProductB{}
}

func (receiver *ProductB) Create() {
	fmt.Println("ProductB:Create")
}

//------------------------------------------------------
//-- ProductAFactory -----------------------------------

type ProductAFactory struct {
}

func NewProductAFactory() ProductFactory {
	return &ProductAFactory{}
}

func (receiver *ProductAFactory) Make() Product {
	return NewProductA()
}

//------------------------------------------------------
//-- ProductBFactory -----------------------------------

type ProductBFactory struct {
}

func NewProductBFactory() ProductFactory {
	return &ProductBFactory{}
}

func (receiver *ProductBFactory) Make() Product {
	return NewProductB()
}

//------------------------------------------------------

/*
 * Client
 */
func main() {
	productAFactory := NewProductAFactory()
	productA := productAFactory.Make()
	productA.Create()

	productBFactory := NewProductBFactory()
	productB := productBFactory.Make()
	productB.Create()
}

/*
ProductA:Create
ProductB:Create
*/
