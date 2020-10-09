package main
import "fmt"
func main (){
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
}
