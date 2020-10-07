package main

import "fmt"

func main() {
	sliOri := []int{108, 107, 105, 109, 103, 102}
	tmpValue := 0
	tmpKey := 0
	for k, v := range sliOri {
		if tmpValue < v {
			tmpValue = v
			tmpKey = k
		}
	}
	sliOri = append(sliOri, tmpValue)
	sliFinal := append(sliOri[:tmpKey], sliOri[tmpKey+1:]...)
	//108, 107, 105, 103, 102, 109
	fmt.Println(sliFinal)

	for i := 0; i < len(sliFinal)-1; i++ {
		if sliFinal[i] > sliFinal[len(sliFinal)-2] {
			sliFinal[i], sliFinal[len(sliFinal)-2] = sliFinal[len(sliFinal)-2], sliFinal[i]
		}
	}
	fmt.Println(sliFinal)
}
