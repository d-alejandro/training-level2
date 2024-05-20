package main

import "fmt"

/*
 Реализовать паттерн «стратегия».
 Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования на практике.
 https://en.wikipedia.org/wiki/Strategy_pattern

 Поведенческий паттерн

 Применимость:
 Инкапсулирует группу родственных алгоритмов, удовлетворяющих одному интерфейсу. На клиента делегирован
 выбор нужного ему алгоритма.

 Плюсы:
 Позволяет избавиться от условных операторов.
 Позволяет заменить наследование за счёт делегирования.
 Реализует принцип открытости/закрытости.
 Скрывает детали реализации алгоритмов от клиента.

 Минусы:
 Увеличивает число объектов в приложении.
 Усложняет использование за счёт изучения особенностей всех алгоритмов.

 Реальное использование:
 В архитектуре Microsoft Windows Driver Framework.
*/

type Strategy interface {
	Execute(int, int)
}

//-- AdditionStrategy ----------------------------------

type AdditionStrategy struct {
}

func NewAdditionStrategy() Strategy {
	return &AdditionStrategy{}
}

func (receiver *AdditionStrategy) Execute(x, y int) {
	fmt.Printf("x + y = %d + %d = %d\n", x, y, x+y)
}

//------------------------------------------------------
//-- SubtractionStrategy -------------------------------

type SubtractionStrategy struct {
}

func NewSubtractionStrategy() Strategy {
	return &SubtractionStrategy{}
}

func (receiver *SubtractionStrategy) Execute(x, y int) {
	fmt.Printf("x - y = %d - %d = %d\n", x, y, x-y)
}

//------------------------------------------------------

type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy}
}

func (receiver *Context) Make(x, y int) {
	receiver.strategy.Execute(x, y)
}

/*
 * Client
 */
func main() {
	additionStrategy := NewAdditionStrategy()
	additionContext := NewContext(additionStrategy)
	additionContext.Make(7, 7)

	subtractionStrategy := NewSubtractionStrategy()
	subtractionContext := NewContext(subtractionStrategy)
	subtractionContext.Make(10, 3)
}

/*
x + y = 7 + 7 = 14
x - y = 10 - 3 = 7
*/
