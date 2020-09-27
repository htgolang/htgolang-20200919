package main

import "fmt"

//不影响其他顺序,只把最大数放到最后
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
	fmt.Println(sliFinal)
}
