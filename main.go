package main

import (
	"fmt"
	_ "test-mnc/docs"
	"test-mnc/routes"
)

// @title           MNC Test API
// @version         1.0
// @description     API untuk test MNC
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  m.iqmal.riffai@gmail.com

// @host      localhost:8082
// @BasePath  /

func main() {
	r := routes.SetupRouter()

	fmt.Println("Server running at http://localhost:8082")
	r.Run(":8082")
}
