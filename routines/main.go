package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	waitGroup.Add(2)

	go func1()
	go func2()

	waitGroup.Wait()
}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func 1:", i)
	}

	waitGroup.Done()
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func 2:", i)
	}
	
	waitGroup.Done()
}
