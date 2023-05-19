package routes

import (
	"log"
	"net/http"

	"github.com/MLCavalcante/go-rest-api/controllers"
	"github.com/MLCavalcante/go-rest-api/middleware"
	"github.com/gorilla/mux"
)

func HandleRequest(){
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/poetas", controllers.GetAllPoetas).Methods("GET")
	r.HandleFunc("/poetas/{id}", controllers.GetPoetaById).Methods("GET")
	r.HandleFunc("/poetas", controllers.CreatePoeta).Methods("POST")
	r.HandleFunc("/poetas/{id}", controllers.DeletePoeta).Methods("DELETE")
	r.HandleFunc("/poetas/{id}", controllers.UpdatePoeta).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}