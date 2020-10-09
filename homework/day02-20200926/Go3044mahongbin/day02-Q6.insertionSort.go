package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main (){

	//6.1
	sli :=  []int{108, 107, 105, 109, 103, 102}
	ls := len(sli)
	for i:=1;i<ls;i++{
		fmt.Println("i,sli[i] : ",i,sli[i])
		tmp := sli[i]

		for j:=i-1;j>=0;j--{
			fmt.Println("j,sli[j] : ",j,sli[j])
			if tmp < sli[j]{
				fmt.Println(tmp,"<",sli[j])
				sli[j+1] = sli[j]
				sli[j] = tmp
				fmt.Println("switch i,j >>>> ",sli)
			}
		}
		fmt.Println("Current Result ------",sli)

	}
	fmt.Println("InsertionSort result:",sli)


	//6.2
	sli2 := []int{102, 103, 105, 107, 108, 109}
	num := rand.New(rand.NewSource(time.Now().UnixNano())).Int() % 10 +100
	fmt.Println("num:", num)

	for i,v := range sli2{
		if v == num {
			fmt.Println("num's index:",i)
			break
		}else{
			if i==len(sli2)-1 {
				fmt.Println("num's index: -1")
			}
		}
	}
}
