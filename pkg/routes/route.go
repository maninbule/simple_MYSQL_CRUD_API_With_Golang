package routes

import (
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/simple_MYSQL_CRUD_API_With_Golang/pkg/controllers"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/book/", controllers.CreateABook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.DeleteABookByID).Methods("DELETE")
	router.HandleFunc("/book", controllers.GetAllBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetABookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateABookByID).Methods("PUT")

	return router
}
