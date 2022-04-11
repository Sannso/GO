package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// go tarea1(1)
	// go tarea1(2)

	// Funcion anonima
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		tarea1(1)
		wg.Done()
	}()

	go func() {
		tarea1(2)
		wg.Done()
	}()

	fmt.Println("Esperando rutinas")
	wg.Wait()
	fmt.Println("Finalizo ejecucion")
}

func tarea1(tarea int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 2)
		fmt.Println("Tarea", tarea)
	}
}
