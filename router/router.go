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

	router.HandleFunc(API_BASE_URL+"/update-task/{id}", middleware.UpdateTask).Methods("PUT", "OPTIONS")
	router.HandleFunc(API_BASE_URL+"/task", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc(API_BASE_URL+"/tasks-by-user", middleware.FetchTasks).Methods("POST", "OPTIONS")
	router.HandleFunc(API_BASE_URL+"/task/delete/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc(API_BASE_URL+"/task/{id}", middleware.TaskStatus).Methods("PUT", "OPTIONS")
	return router

}
