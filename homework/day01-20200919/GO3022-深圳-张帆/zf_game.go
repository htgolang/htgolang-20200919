package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main(){
	rand.Seed(time.Now().Unix())
	//
	sj :=rand.Int() %100    //产生一个0~100的随即数
	for a :=1;a < 10;a++{
		if a <= 5{
			var  number int
			fmt.Print("猜数字游戏0~100有五次机会，请输入你猜的结果:_")
			fmt.Scan(&number)
			if number < sj {
				fmt.Println("你猜小了，请继续猜")
			}else if number > sj {
				fmt.Println("你猜大了，请继续猜")
			} else if  number == sj {
				fmt.Println("小可爱，恭喜你猜对了")
				break     //猜对了，退出循环
			}
		}
		if a > 5{
			fmt.Println("笨蛋，你五次机会已经用完了")
			break     //大于五次机会 退出游戏
		}
	}
}