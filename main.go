package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

// Defino los tipos de Datos.
type Numeros struct {
	Valores []int
}

// Creamos una Instancia
func main() {
	text := "{\"Valores\": [4, 5, 2, 3, 7, 8, 9, 1, 10, 11, 56, 0, 13]}"
	bytes := []byte(text)
	var resultados Numeros
	json.Unmarshal(bytes, &resultados)
	fmt.Printf("Sin Clasificar = %v", resultados.Valores)

	/*El paquete `sort` de Go implementa ordenamiento para los tipos primitivos y definidos por el usuario.*/
	ints := []int{4, 5, 2, 3, 7, 8, 9, 1, 10, 11, 56, 0, 13}
	sort.Ints(ints)
	crear_json, _ := json.Marshal(ints)
	convertir_a_cadena := string(crear_json)
	fmt.Println(" Clasificado = ", convertir_a_cadena)
	// fmt.Println(" Clasificado = ", ints)
}
