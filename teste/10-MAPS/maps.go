package main

import "fmt"

func main() {
	fmt.Println("maps")

	usuario := map[string]map[string]string{
		"nome": {
			"primeiro":  "cesar",
			"sobrenome": "gomes",
		},
		"curso": {
			"nome do curso": "Engenharia da computação",
			"faculdade":     "anhanguera",
		},
	}
	fmt.Println(usuario)
	delete(usuario, "nome")
	fmt.Println(usuario)
}
