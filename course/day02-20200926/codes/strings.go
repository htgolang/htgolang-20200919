package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("1", "2"))
	fmt.Println(strings.Contains("abc", "bd"))
	fmt.Println(strings.Count("abcabdabd", "bd"))
	fmt.Printf("%q\n", strings.Fields("abc abc\t123\n1\r1"))
	fmt.Println(strings.HasPrefix("abcd", "ab"))
	fmt.Println(strings.HasSuffix("abcd", "ab"))

	fmt.Println(strings.Join([]string{"1", "2"}, "-"))
	fmt.Println(strings.Repeat("*", 20))
	fmt.Println(strings.Replace("abcabcabcdabc", "abc", "x", 2))
	fmt.Println(strings.ReplaceAll("abcabcabcdabc", "abc", "x"))
	fmt.Printf("%q\n", strings.Split("a-b--c", "-"))
	fmt.Printf("%q\n", strings.SplitAfter("a-b-c", "-"))
	fmt.Printf("%q\n", strings.SplitN("a-b-c", "-", 2))
	fmt.Println(strings.Title("a b bab abc"))
	fmt.Println(strings.ToLower("A B bab abc"))
	fmt.Println(strings.ToUpper("A B bab abc"))
	fmt.Println(strings.Trim("abcbcdabcdbcabc", "abc"))
	fmt.Println(strings.TrimLeft("abcbcdabcdbcabc", "abc"))
	fmt.Println(strings.TrimRight("abcbcdabcdbcabc", "abc"))
	fmt.Println(strings.TrimSpace(" abcbcdabcdbcabc \n\r   "))
	fmt.Println(strings.TrimSuffix("abcbcdabcdbcabc", "abc"))
	fmt.Println(strings.TrimPrefix("abcbcdabcdbcabc", "abc"))
}
