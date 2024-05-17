package main

import "fmt"

/*
 Реализовать паттерн «строитель».
 Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования на практике.
 https://en.wikipedia.org/wiki/Builder_pattern

 Применимость:
 Позволяет создавать сложные объекты за счёт поэтапного разбиения процесса конструирования. Код Builder-а состоит из
 множества методов, каждый из которых соответствует определённому этапу построения объекта. Вызовы этих методов
 делегированы клиенту (существуют варианты паттерна, где эта обязанность принадлежит отдельному посреднику - директору).
 Клиенты комбинируют вызовы методов в зависимости от необходимости, и таким образом создаются различные представления
 объектов одного и того же кода Builder-а.

 Плюсы:
 Простота использования клиентом.
 Простота проектирования архитектуры Builder-а.
 Возможность создавать различные представления объектов.
 Даёт более тонкий контроль над процессом конструирования.

 Минусы:
 Высокая степень связанности между Builder-ом и клиентом.
 Необходимо изучать алгоритм создания объекта из-за определённого порядка вызова методов.
 При повторном создании объектов с одинаковыми комбинациями вызовов методов необходимо дублировать код. То есть нет
 возможности объединить вызовы методов в одну группу и в дальнейшем применять её к вновь созданному объекту.

 Реальное использование:
 Создание построителя запросов к БД в ORM-библиотеках.
*/

type Builder interface {
	BuildStepA() Builder
	BuildStepB() Builder
	GetText() string
}

type BuilderImplementation struct {
	text string
}

func NewBuilderImplementation() Builder {
	return &BuilderImplementation{}
}

func (receiver *BuilderImplementation) BuildStepA() Builder {
	receiver.setText("BuilderImplementation:BuildStepA")

	return receiver
}
func (receiver *BuilderImplementation) BuildStepB() Builder {
	receiver.setText("BuilderImplementation:BuildStepB")

	return receiver
}
func (receiver *BuilderImplementation) GetText() string {
	return receiver.text
}

func (receiver *BuilderImplementation) setText(text string) {
	if len(receiver.text) == 0 {
		receiver.text += text
		return
	}

	receiver.text += " <-> " + text
}

/*
 * Client
 */
func main() {
	builder := NewBuilderImplementation()

	text := builder.
		BuildStepA().
		BuildStepB().
		GetText()

	fmt.Println("\n", text)
}

/*
 BuilderImplementation:BuildStepA <-> BuilderImplementation:BuildStepB
*/
