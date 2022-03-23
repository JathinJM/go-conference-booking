package main

import (
	"fmt"
)

func message(country string) string {
	return "welome to " + country + "!"
}

func switch_flow() {
	var country string

	fmt.Printf("please enter the country:\n")
	fmt.Scan(&country)

	switch country {
	case "India":
		fmt.Printf(message(country))

	case "England":
		fmt.Printf(message(country))

	case "Canada":
		fmt.Printf(message(country))

	case "America":
		fmt.Printf(message(country))

	case "Germany":
		fmt.Printf(message(country))

	default:
		fmt.Printf("Invalid country")
	}
}
