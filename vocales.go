package Ejercicios

import (
	"fmt"
	"strings"
)

// La función ContarVocales cuenta las vocales en una frase y al final imprime el conteo.

func ContarVocales(texto string) {

	texto = strings.ToLower(texto) // Aquí se pasa todo el texto a minúsculas

	var a, e, i, o, u int // Aquí inicializo contadores para cada vocal.

	for _, letra := range texto {
		switch letra {
		case 'a':
			a++
		case 'e':
			e++
		case 'i':
			i++
		case 'o':
			o++
		case 'u':
			u++
		}
	}

	fmt.Println("Recuento de vocales:")
	fmt.Printf("a: %d, e: %d, i: %d, o: %d, u: %d\n", a, e, i, o, u)
}
