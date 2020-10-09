package main

import (
	"fmt"
	"strings"
)

func main() {
	const ss string = `I am happy to join with you today in what will go down in history as the greatest demonstration for freedom in the history of our nation.
Five score years ago, a great American, in whose symbolic shadow we stand today, signed the Emancipation Proclamation. This momentous decree came as a great beacon light of hope to millions of Negro slaves who had been seared in the flames of withering injustice. It came as a joyous daybreak to end the long night of their captivity.
But one hundred years later, the Negro still is not free. One hundred years later, the life of the Negro is still sadly crippled by the manacles of segregation and the chains of discrimination. One hundred years later, the Negro lives on a lonely island of poverty in the midst of a vast ocean of material prosperity. One hundred years later, the Negro is still languished in the corners of American society and finds himself an exile in his own land. And so we've come here today to dramatize a shameful condition.`
	// Convert ss to lower rune
	ssLow := []rune(strings.ToLower(ss))
	//fmt.Println(ss_low)
	static := map[string]int{}
	for _, v := range ssLow {
		// Limit the statistical scope from 'a' to 'z'
		if v >= 'a' && v <= 'z' {
			static[string(v)]++
		}
	}
	fmt.Println(static)
}
