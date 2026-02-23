package main

import "fmt"

func main() {
	carro := []int{23452, 72622, 5433, 9273, 62533, 62353, 15625, 25162, 41312, 62762, 62532, 52426}
	primerTrimestre := carro[:3]

	fmt.Println(primerTrimestre)
	total := 0

	for _,v := range primerTrimestre{
		total += v

	}
	fmt.Println(total)
	//Tenint en compte que la col.lecció de dades que mostrem, correspon a la facturació d'una tenda de cada mes del any, tens que mostrar el total facturat el primer trimestre.

}
