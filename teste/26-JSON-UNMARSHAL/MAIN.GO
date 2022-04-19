package main

import ( "fmt"
"encoding/json"
"log"
)

type cachorro struct{
	Nome string  `json:"-"`
	Raca string  `json:"raca"`
	Idade uint `json:"idade"`
}

func main(){
	cachorroEmJson := `{"nome":"Rex","raca":"pitbull","idade":3}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJson), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)
	
	cachorro2EmJson := `{"nome":"Leo","raca":"Dalmata"}`

	c2 := make (map[string]string)
	if erro := json.Unmarshal([]byte(cachorro2EmJson),&c2); erro != nil{
		log.Fatal(erro)
	}
	fmt.Println(c2)
}