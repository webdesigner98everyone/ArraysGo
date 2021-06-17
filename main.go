package main

import "fmt"

func ordenarMayor(listNum []int, Cant int) {
	tmp := 0
	for x := 0; x < Cant; x++ {
		for y := 0; y < Cant; y++ {
			if listNum[x] > listNum[y] {
				tmp = listNum[x]
				listNum[x] = listNum[y]
				listNum[y] = tmp
			}
		}
	}
	fmt.Print("\nArray dinamico ordenado: ")
	for i := 0; i < Cant; i++ {
		fmt.Print("[", listNum[i], "]")
	}
	fmt.Println()
}
func ordenarMenor(listNum []int, Cant int) {
	tmp := 0
	for x := 0; x < Cant; x++ {
		for y := 0; y < Cant; y++ {
			if listNum[x] < listNum[y] {
				tmp = listNum[y]
				listNum[y] = listNum[x]
				listNum[x] = tmp
			}
		}
	}
	fmt.Print("\nArray dinamico ordenado: ")
	for i := 0; i < Cant; i++ {
		fmt.Print("[", listNum[i], "]")
	}
	fmt.Println()
}
func main() {
	listNum := make([]int, 0)
	var cant, opt, valor int
	fmt.Print("¿Cuantos valores desea ingresar?: ")
	fmt.Scanf("%d", &cant)
	for v := 1; v <= cant; v++ {
		fmt.Print("Ingrese en valor #", v, ": ")
		fmt.Scanf("%d", &valor)
		listNum = append(listNum, valor)
	}
	fmt.Println("\nSelecciones una de las siguientes opciones")
	fmt.Println("------------------------------------------")
	fmt.Println("1) Ordenar de mayor a menor")
	fmt.Println("2) Ordenar de menor a mayor")
	fmt.Println("------------------------------------------")
	fmt.Print("Opcion: ")
	fmt.Scanf("%d", &opt)
	switch opt {
	case 1:
		ordenarMayor(listNum, cant)
	case 2:
		ordenarMenor(listNum, cant)
	default:
		fmt.Println("Opcion no valida!!")
	}
}
