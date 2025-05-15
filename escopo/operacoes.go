package main

import "fmt" 

type Carro struct{
	modelo string
	ano string
	velocidade string
}

func exibeDados(c Carro){
fmt.Println("Modelo do carro; ", c.modelo)
fmt.Println("Ano do carro; ", c.ano)
fmt.Println("Velocidade do carro; ", c.velocidade)
}

func main() {
	c1 := Carro{"HB20", 2019, 190 km}
	exibeDados(c1)
}


