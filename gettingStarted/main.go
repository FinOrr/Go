package main

import (
	"example.com/gettingStarted/callingExternalPackages"
	"fmt"
)

func main() {
	fmt.Println(callingExternalPackages.callQuotePackage())
}
