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
	text := "{\"Valores\": [3, 5, 5, 6, 8, 3, 4, 4, 7, 7, 1, 1, 2]}"
	bytes := []byte(text)
	var resultados Numeros
	json.Unmarshal(bytes, &resultados)
	fmt.Printf("Sin Clasificar = %v", resultados.Valores)

	/*El paquete `sort` de Go implementa ordenamiento para los tipos primitivos y definidos por el usuario.*/
	ints := []int{3, 5, 5, 6, 8, 3, 4, 4, 7, 7, 1, 1, 2}
	sort.Ints(ints)
	crear_json, _ := json.Marshal(ints)
	convertir_a_cadena := string(crear_json)
	fmt.Println(" Clasificado = ", convertir_a_cadena)
	// fmt.Println(" Clasificado = ", ints)
}
