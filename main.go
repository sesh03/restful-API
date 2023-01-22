package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Task struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Completed   bool   `json:"completed"`
}

var tasks []Task

func main() {
    http.HandleFunc("/tasks", handleTasks)
    http.HandleFunc("/tasks/", handleTask)
    http.ListenAndServe(":8080", nil)
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        json.NewEncoder(w).Encode(tasks)
    case "POST":
        var task Task
        json.NewDecoder(r.Body).Decode(&task)
        task.ID = len(tasks) + 1
        tasks = append(tasks, task)
        json.NewEncoder(w).Encode(task)
    }
}

func handleTask(w http.ResponseWriter, r *http.Request) {
    taskID := r.URL.Path[len("/tasks/"):]
    switch r.Method {
    case "GET":
        for _, task := range tasks {
            if strconv.Itoa(task.ID) == taskID {
                json.NewEncoder(w).Encode(task)
                return
            }
        }
        http.Error(w, "Task not found", http.StatusNotFound)
    case "PUT":
        var task Task
        json.NewDecoder(r.Body).Decode(&task)
        for i, t := range tasks {
            if strconv.Itoa(t.ID) == taskID {
                task.ID, tasks[i] = tasks[i].ID, task
                json.NewEncoder(w).Encode(task)
                return
            }
        }
        http.Error(w, "Task not found", http.StatusNotFound)
    case "DELETE":
        for i, task := range tasks {
            if strconv.Itoa(task.ID) == taskID {
                tasks = append(tasks[:i], tasks[i+1:]...)
                json.NewEncoder(w).Encode(task)
                return
            }
        }
        http.Error(w, "Task not found", http.StatusNotFound)
    }
}
