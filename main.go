package main

import (
	"fmt"
	"main/models"
	"main/routes"
	"net/http"
)

func main() {

	models.Post{}.Migrate()
	fmt.Println("Server : https://localhost:8080")
	http.ListenAndServe(":8080", routes.Routes())
}
