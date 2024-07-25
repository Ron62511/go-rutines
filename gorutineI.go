package main

import "fmt"

func main() {

	fmt.Println("Esto es una gorutine")
	foo()

}

func foo() {

	for i := 0; i <= 10; i++ {
		fmt.Println("foo()", i)
	}
}
