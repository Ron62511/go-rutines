package main

import (
	"fmt"
	"runtime"
	"sync"
)

//Paquete sync
// Aplicando WaitGroups()

var wg sync.WaitGroup

func main() {

	fmt.Println("Esto es una gorutine")
	runtimePack()

	wg.Add(1)
	go foo()
	bar()
	wg.Wait()

}

func foo() {

	for i := 0; i <= 10; i++ {
		fmt.Println("foo()", i)
	}
	wg.Done()
}

func bar() {

	for j := 0; j <= 10; j++ {

		fmt.Println("bar()", j)
	}
}
func runtimePack() {
	//Paquete runtime
	fmt.Println("OS", runtime.GOOS)                   //Sistema Operativo
	fmt.Println("ARCH", runtime.GOARCH)               //Arquitectura
	fmt.Println("CPU", runtime.NumCPU())              //Numero de CPU
	fmt.Println("#Gorutines", runtime.NumGoroutine()) //Numero de Gorutines

}
