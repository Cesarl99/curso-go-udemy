package main

import ("fmt"
"encoding/json"
"log"
"bytes")

type cachorro struct{
	Nome string  `json:"nome"`
	Raca string  `json:"raca"`
	Idade uint `json:"idade"`
	
}
func main(){
	c:= cachorro{"REX","DALMATA",5}
	fmt.Println(c)

	cachorroEmJson,erro:=json.Marshal(c)
	if erro != nil{
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJson)

	fmt.Println(bytes.NewBuffer(cachorroEmJson))

	c2 := map[string]string{
		"nome":"leo",
		"raca":"pitbull",
	}
	cachorro2Emjson,erro := json.Marshal(c2)
	if erro != nil{
		log.Fatal(erro)
	}
	fmt.Println(cachorro2Emjson)
	fmt.Println(bytes.NewBuffer(cachorro2Emjson))
}