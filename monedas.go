package Ejercicios

import (
	"fmt"
	"strings"
)

// La función ConvertirDolares convierte un valor dado en dólares a otra moneda específica

// Esta función debe recibir un valor a convertir y luego la moneda (o sea escribiendo como string
// el nombre de la moneda)a la que se desea hacer el cambio.

func ConvertirDolares(valorDolares float64, monedaNueva string) {

	monedaNueva = strings.ToLower(monedaNueva) //Aquí se hace un "conversion" o transformación para que el programa lea tanto min como may

	var valorConvertido float64 // Aquí declaro la variable que guarda el valor que se va a convertir

	switch monedaNueva {

	case "euros":
		valorConvertido = valorDolares * 0.85
		fmt.Printf("%.2f dólares equivalen a %.2f euros.\n", valorDolares, valorConvertido)

	case "libras":
		valorConvertido = valorDolares * 0.74
		fmt.Printf("%.2f dólares equivalen a %.2f libras esterlinas.\n", valorDolares, valorConvertido)

	case "won":
		valorConvertido = valorDolares * 1401.0
		fmt.Printf("%.2f dólares equivalen a %.2f wones surcoreanos.\n", valorDolares, valorConvertido)

	case "btc":
		valorConvertido = valorDolares * 0.00000895
		fmt.Printf("%.2f dólares equivalen a %.6f BTC.\n", valorDolares, valorConvertido)

	default:
		fmt.Println("Error: La moneda de destino no es válida.")
	}
}
