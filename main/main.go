package main

import (
  "fmt"
  "github.com/ariauakbar/todos/models"
  "path"
  "html/template"
  "net/http"
  // "log"
)

func main() {

http.HandleFunc("/", IndexHandler)
http.ListenAndServe(":8080", nil)

  //models.CreateTodoTable()

  // for i := 0; i < 5; i++{
  //   models.CreateTodo("Creating Presentation")
  // }
  todos := models.GetAllTodos()
  for _, todo := range todos {
    fmt.Println(todo.Name)
  }
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
  }
}
