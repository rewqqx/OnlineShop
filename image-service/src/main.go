package main

import (
	"backend/src/utils"
	"fmt"
)

func main() {
	fmt.Println("Start Service:")
	server := utils.New()
	server.Start(8081)
}
