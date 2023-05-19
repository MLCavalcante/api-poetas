package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	

	"github.com/MLCavalcante/go-rest-api/database"
	"github.com/MLCavalcante/go-rest-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAllPoetas(w http.ResponseWriter, r *http.Request){
	var p []models.Poeta
	database.DB.Table("poetas").Order("id ASC").Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetPoetaById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Retorna um mapa com os parâmetros da requisição
	id := vars["id"] // Acessa o parâmetro id
    var p []models.Poeta // Cria uma variável do tipo Poeta
	database.DB.Table("poetas").First(&p, id) // Busca o registro no banco de dados
	json.NewEncoder(w).Encode(p) // Codifica o JSON e envia para o cliente
	
}

func CreatePoeta(w http.ResponseWriter, r *http.Request){
	var novopoeta models.Poeta  // Cria uma variável do tipo Poeta
	json.NewDecoder(r.Body).Decode(&novopoeta)// Decodifica o JSON recebido e armazena na variável novopoeta
	database.DB.Table("poetas").Create(&novopoeta) // Cria um novo registro no banco de dados
	json.NewEncoder(w).Encode(novopoeta) // Codifica o JSON e envia para o cliente
}



func DeletePoeta(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r) // Retorna um mapa com os parâmetros da requisição
	id := vars["id"] // Acessa o parâmetro id
	var p []models.Poeta // Cria uma variável do tipo Poeta
	database.DB.Table("poetas").Delete(&p, id) // Deleta o registro no banco de dados
	json.NewEncoder(w).Encode(p) // Codifica o JSON e envia para o cliente

}

func UpdatePoeta(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r) // Retorna um mapa com os parâmetros da requisição
	id := vars["id"] // Acessa o parâmetro id
	var p []models.Poeta // Cria uma variável do tipo Poeta
	database.DB.Table("poetas").First(&p, id) // Busca o registro no banco de dados
	json.NewDecoder(r.Body).Decode(&p) // Decodifica o JSON recebido e armazena na variável p
	database.DB.Table("poetas").Save(&p) // Atualiza o registro no banco de dados
	json.NewEncoder(w).Encode(p) // Codifica o JSON e envia para o cliente
}

