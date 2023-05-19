package main

import (
	"fmt"

	"github.com/MLCavalcante/go-rest-api/database"
	"github.com/MLCavalcante/go-rest-api/models"
	"github.com/MLCavalcante/go-rest-api/routes"
)



func main() {
	models.Poetas = []models.Poeta{
		{Id:1, Nome:"nome 1", Historia:"historia 1"},
		{Id:2, Nome:"nome 2", Historia:"historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor...")
	routes.HandleRequest()
}