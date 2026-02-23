package main

import "fmt"

func main() {
	carro := []string{"pastanagues", "pebrots", "aigua", "formatge"}
	carro[0] = "cogombre"

	fmt.Println(carro)
	//Sobre el carro de la compra que et mostrem, afegeix un cogombre en la posició 1 substituint l'altre article.
}
