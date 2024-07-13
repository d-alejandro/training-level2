Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Программа выведет error. Функция test() вернёт указатель на нулевую структуру customError. При присвоении переменной
err произойдёт приведение указателя нулевой структуры customError к интерфейсу error. Следовательно, при сравнении err
не будет равна nil и выражение err != nil будет верно. 
```
