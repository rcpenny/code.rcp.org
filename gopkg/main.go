package main

import (
	"fmt"

	"code.rcp.org/gopkg/rc_strings"
)

func main() {
	brands := []string{"Apple", "Google", "Microsoft"}
	allBrands := rc_strings.Join(brands, ",")
	fmt.Println(allBrands)
}
