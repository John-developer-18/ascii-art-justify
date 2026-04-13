package utils

import (
	"fmt"
	"slices"
)

var validOptions, validBanners []string = []string{"left", "right", "center", "justify"}, []string{"shadow", "tinkertoy", "standard"}

func receiveInput(option string, strVal string, banner string) (string, error) {

	//Processing options and banners
	if option == "" || banner == "" {
		option = "left"
	}
	if !slices.Contains(validOptions, option) || !slices.Contains(validBanners, banner) {
		return "", fmt.Errorf("Invalid")
	}

	//

	return "", nil

}
