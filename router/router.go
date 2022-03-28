package router

import (
	"golang-react-to-do/server/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	API_BASE_URL := "/go-api"
	router := mux.NewRouter()
	router.HandleFunc(API_BASE_URL+"/user/add", middleware.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc(API_BASE_URL+"/user/login", middleware.LoginUser).Methods("POST", "OPTIONS")

	router.HandleFunc(API_BASE_URL+"/note/{id}", middleware.UpdateTask).Methods("PUT", "OPTIONS")
	router.HandleFunc(API_BASE_URL+"/note", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc(API_BASE_URL+"/notes-by-user", middleware.FetchTasks).Methods("POST", "OPTIONS")
	router.HandleFunc(API_BASE_URL+"/note/delete/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc(API_BASE_URL+"/note/{id}", middleware.TaskStatus).Methods("PUT", "OPTIONS")
	return router

}
