package main

func generica (interf interface{}){
	fmt.Println(interf)
}

func main(){
	generica("nome")
	generica(1)
	generica(true)
	generica(1.5)
}