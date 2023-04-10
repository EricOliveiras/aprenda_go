package main

import "fmt"

func main() {
	sayHello("Eric")
}

func sayHello(name string)  {
	fmt.Println("Ola,", name, "!")
}