package middleware

import (
	"encoding/json"
	"fmt"
	"golang-react-to-do/server/database"
	"golang-react-to-do/server/models"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		return
	}
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	fmt.Println("create task hit with task ", task)
	if task.Message == "" {
		return
	}
	var createdTask = database.InsertTask(task)
	json.NewEncoder(w).Encode(createdTask)
}
func FetchTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println(r.Method)
	if r.Method == "OPTIONS" {
		return
	}
	var user models.User
	fmt.Println(user)
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("fetch tasks hit with user", user)
	var tasksList = database.GetTasksByUser(user)
	if tasksList == nil {
		tasksList = make([]models.Task, 0)
	}
	json.NewEncoder(w).Encode(tasksList)
}
func DeleteTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		return
	}
	fmt.Println("delete hit")

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("id not found")
	}
	fmt.Println(`ObjectID := `, id)
	database.DeleteTask(id)
	json.NewEncoder(w).Encode("Task Deleted Successfully")
}
func TaskStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Done task hit")
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}
	fmt.Println(`id := `, id)
	database.TaskStatus(id)
}
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Println("In update")
	if r.Method == "OPTIONS" {
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Println(`update id := `, id)
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	fmt.Println("update task hit with task ", task)

	database.UpdateTask(id, task)
}
