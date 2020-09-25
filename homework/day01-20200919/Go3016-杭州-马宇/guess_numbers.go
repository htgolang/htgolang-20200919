package main
import (
	"fmt"
	"math/rand"
)
func main() {
	random := rand.Intn(100)
	for {
		for i := 0; i < 5; i++ {
			var test int
			fmt.Println("猜数字: ")
			fmt.Scan(&test)
			for {
				if random == test {
					goto START

				} else if test > random {
					fmt.Println("对不起，猜太大了")
					break
				} else if test < random  {
					fmt.Println("对不起，猜太小了")
					break
				}
			}
		}
		fmt.Println("未猜对，你太笨了")
		break
	START:
		fmt.Println("太聪明了")
		break
	}
}
