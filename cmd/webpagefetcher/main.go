package main

import (
	"fmt"

	"WebPageFetcher/service"
)

func main() {
	fmt.Println("Please see README.md for usage. Enter commands: ")
	service.FetchWebpages()
	fmt.Println("Exiting...")
}
