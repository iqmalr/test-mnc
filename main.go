package main

import (
	"fmt"
	"test-mnc/routes"
)

func main() {
	r := routes.SetupRouter()

	fmt.Println("Server running at http://localhost:8081")
	r.Run(":8081")
}
