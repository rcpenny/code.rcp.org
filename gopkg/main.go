package main

import (
	"fmt"

	"code.rcp.org/gopkg/src/rc_strings"
)

func main() {
	brands := []string{"Apple", "Google", "Microsoft"}
	allBrands := rc_strings.Join(brands, ",")
	fmt.Println(allBrands)
}
