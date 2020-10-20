package funcs

import (
	"fmt"

	"moul.io/banner"
)

//CMSBanner ...
// Generate CMS banner
func CMSBanner() {
	fmt.Println("---------------------")
	fmt.Println(banner.Inline("c m s"))
	fmt.Println("---------------------")
}
