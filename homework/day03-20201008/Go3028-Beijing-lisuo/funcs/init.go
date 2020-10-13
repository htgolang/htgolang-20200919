package funcs

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
}
func GenId() {
	// gen a number in [)
	result, _ := rand.Int(rand.Reader, big.NewInt(999999999999))
	fmt.Println(result)
}
func Init() {
	GenId()
}
