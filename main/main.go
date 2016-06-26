package main

import (
  //"fmt"
  "github.com/ariauakbar/todos/models"
  "path"
  "html/template"
  "net/http"
  "strconv"
  // "log"
)

func main() {
  http.HandleFunc("/", IndexHandler)
  http.ListenAndServe(":8080", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request)  {

  if r.Method == "GET" {
    path := path.Join("templates", "index.html")
    tmpl, err := template.ParseFiles(path)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    todos := models.GetAllTodos()
    if err := tmpl.Execute(w, todos); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  } else if r.Method == "POST" {
    name := r.FormValue("name")
    models.CreateTodo(name)
    http.Redirect(w, r, r.Referer(), http.StatusMovedPermanently)
  } else if r.Method == "PUT" {
    todoID := r.FormValue("id")
    id, err := strconv.Atoi(todoID)
    if err != nil {
      panic(err.Error())
    }
    models.ToggleFinishedByID(id)
  }
}
