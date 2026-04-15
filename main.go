package main

import (
	"ascii/justify/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	errorFlag := `
Usage: go run . [OPTION] [STRING] [BANNER]

Example: go run . --align=right something standard
	`
	option, banner, strVal := "", "", ""

	switch len(os.Args) {
	case 2:
		strVal = os.Args[1]
		str, err := utils.ReceiveInput(option, strVal, banner)

		if err != nil {
			log.Fatal("Error", err)
		}

		fmt.Println(str)
	case 3:
		strVal = os.Args[2]
		option = strings.TrimSuffix(os.Args[1], "--align==")
		str, err := utils.ReceiveInput(option, strVal, banner)

		if err != nil {
			log.Fatal("Error", err)
		}

		fmt.Println(str)
	case 4:
		strVal = os.Args[2]
		option = strings.TrimSuffix(os.Args[1], "--align=")
		banner = os.Args[3]
		str, err := utils.ReceiveInput(option, strVal, banner)

		if err != nil {
			log.Fatal("Error", err)
		}

		fmt.Println(str)
	default:
		fmt.Println(errorFlag)
	}

	fmt.Println(errorFlag, option, banner, strVal)

}
