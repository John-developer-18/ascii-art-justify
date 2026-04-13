package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	errorFlag := `
Usage: go run . [OPTION] [STRING] [BANNER]

Example: go run . --align=right something standard
	`
	option, banner, strVal := "", "", ""

	args := os.Args[1:4]
    
	switch args{
	case 1:

	case 2:

	case 3:

	default:
	   fmt.Println(errorFlag)
	}

	fmt.Println(errorFlag, option, banner, strVal)

}


