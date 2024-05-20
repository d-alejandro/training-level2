package main

import "fmt"

/*
 Реализовать паттерн «команда».
 Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования на практике.
 https://en.wikipedia.org/wiki/Command_pattern

 Поведенческий паттерн

 Применимость:
 Инкапсулирует операции в отдельные объекты, позволяя передавать их другим функциям и объектам
 в качестве параметра, ставить операции в очередь, логировать, а также поддерживать их отмену.

 Плюсы:
 Придает системе гибкость, отделяя инициатора операции (Client) от его получателя (Receiver).
 Реализует принцип открытости/закрытости.
 Упрощает работу с транзакциями СУБД.
 Упрощает работу с системами управления событиями.

 Минусы:
 Усложняет код программы из-за введения дополнительных программных сущностей.

 Реальное использование:
 В Java-библиотеке Swing и Borland Delphi объект Action является объектом команды.
*/

//-- Receiver ------------------------------------------

type Receiver struct {
}

func NewReceiver() *Receiver {
	return &Receiver{}
}

func (receiver *Receiver) Action() {
	fmt.Println("Receiver:Action")
}

//------------------------------------------------------

type Command interface {
	Execute()
}

//-- BeforeCommand -------------------------------------

type BeforeCommand struct {
	*Receiver
}

func NewBeforeCommand(receiver *Receiver) Command {
	return &BeforeCommand{receiver}
}

func (receiver *BeforeCommand) Execute() {
	fmt.Println("BeforeCommand:Execute")
	receiver.Action()
}

//------------------------------------------------------
//-- AfterCommand --------------------------------------

type AfterCommand struct {
	*Receiver
}

func NewAfterCommand(receiver *Receiver) Command {
	return &AfterCommand{receiver}
}

func (receiver *AfterCommand) Execute() {
	receiver.Action()
	fmt.Println("AfterCommand:Execute")
}

//------------------------------------------------------
//-- Invoker -------------------------------------------

type Invoker struct {
	command Command
}

func NewInvoker(command Command) *Invoker {
	return &Invoker{command}
}

func (receiver *Invoker) Execute() {
	receiver.command.Execute()
}

//------------------------------------------------------

/*
 * Client
 */
func main() {
	receiver := NewReceiver()

	beforeCommand := NewBeforeCommand(receiver)
	beforeInvoker := NewInvoker(beforeCommand)
	beforeInvoker.Execute()

	afterCommand := NewAfterCommand(receiver)
	afterInvoker := NewInvoker(afterCommand)
	afterInvoker.Execute()
}

/*
BeforeCommand:Execute
Receiver:Action
Receiver:Action
AfterCommand:Execute
*/
