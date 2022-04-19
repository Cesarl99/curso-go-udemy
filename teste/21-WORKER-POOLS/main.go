package main

import (
	"fmt"
)

func main() {
	x := 77
	tarefas := make(chan int, x)
	resultados := make(chan int, x)

	go worker(tarefas, resultados)

	for i := 0; i < x; i++ {
		tarefas <- i
	}
	close(tarefas)
	for i := 0; i < x; i++ {
		resultados := <-resultados
		fmt.Println(resultados)
	}
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}
func fibonacci(posição int) int {
	if posição <= 1 {
		return posição
	}
	return fibonacci(posição-2) + fibonacci(posição-1)
}
