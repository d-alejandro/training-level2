package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих
каналов закроется. Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало
эту связь, однако иногда неизвестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов,
реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“done after %v”, time.Since(start))
*/

/*
Output:
done after 1.000331845s
*/
func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}

	or = mergeChannels

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})

		go func() {
			defer close(c)
			time.Sleep(after)
		}()

		return c
	}

	start := time.Now()

	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start))
}

func mergeChannels(inputChannels ...<-chan any) <-chan any {
	var (
		waitGroup     sync.WaitGroup
		mutex         sync.Mutex
		quitFlag      bool
		outputChannel = make(chan any)
	)

	waitGroup.Add(1)

	for _, inputChannel := range inputChannels {
		go func(inputChan <-chan any) {
			for {
				if _, isClose := <-inputChan; isClose == false {
					mutex.Lock()
					if quitFlag == false {
						quitFlag = true
						waitGroup.Done()
					}
					mutex.Unlock()
				}
			}
		}(inputChannel)
	}

	go func() {
		defer close(outputChannel)
		waitGroup.Wait()
	}()

	return outputChannel
}
