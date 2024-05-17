package main

import "fmt"

/*
 Реализовать паттерн «фасад».
 Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования на практике.
 https://en.wikipedia.org/wiki/Facade_pattern

 Структурный паттерн

 Применимость:
 Решение проблемы группировки классов после применения Single Responsibility Principle.
 Adapter и Facade являются "обёртками". Цель Facade – создание нового более простого интерфейса,
 цель Adapter – адаптация существующего интерфейса. Facade обычно "обёртывает" несколько объектов,
 а Adapter – один объект.

 Плюсы:
 Упрощает использование подсистемы за счёт уменьшения времени на её изучение.
 Способствует уменьшению степени связанности между подсистемой и клиентом.
 Инкапсулирует подсистему.
 Уменьшает компиляционные зависимости в больших программных системах.

 Минусы:
 Для клиентов компоненты подсистемы остаются чёрным ящиком.

 Реальное использование:
 На базе данного паттерна реализован доступ к базовым классам PHP-фреймворка Laravel.
*/

type SubsystemA struct {
}

func NewSubsystemA() *SubsystemA {
	return &SubsystemA{}
}

func (receiver *SubsystemA) Make() string {
	return "SubsystemA:Make"
}

type SubsystemB struct {
}

func NewSubsystemB() *SubsystemB {
	return &SubsystemB{}
}

func (receiver *SubsystemB) Build() string {
	return "SubsystemB:Build"
}

type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
}

func NewFacade() *Facade {
	return &Facade{
		subsystemA: NewSubsystemA(),
		subsystemB: NewSubsystemB(),
	}
}

func (receiver *Facade) Execute() string {
	return receiver.subsystemA.Make() + " <-> " + receiver.subsystemB.Build()
}

/*
 * Client
 */
func main() {
	facade := NewFacade()

	fmt.Println("\n", facade.Execute())
}

/*
 SubsystemA:Make <-> SubsystemB:Build
*/
