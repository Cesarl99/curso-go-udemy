package endereços

import "strings"

//TipodeEndereços
func TipodeEndereços (endereço string)string{
	TiposValidos := []string{"rua","avenida","estrada","rodovia"}
	endereçoEmLetraMinuscula := strings.ToLower(endereço)
	primeiraPalavraDeEndereço := strings.Split(endereçoEmLetraMinuscula," ")[0]	
	//strings.splits divide a frase em slice 

	endereçosTemUmTipoValido := false
	for _ ,tipo := range TiposValidos{
		if tipo == primeiraPalavraDeEndereço{
			endereçosTemUmTipoValido == true
		}
	}
	if endereçosTemUmTipoValido{
		return primeiraPalavraDeEndereço
	}
	return "tipo invalido"
}