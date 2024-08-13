package main

import (
	"github.com/Bruno-Bianchi/api-go-gin/database"
	"github.com/Bruno-Bianchi/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
