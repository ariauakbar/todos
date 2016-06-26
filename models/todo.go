package models

import (
  "github.com/jinzhu/gorm"
)

var gdb, _ = GetDBConnection()

type Todo struct {
  gorm.Model
  Name string `sql:"type:varchar(200)" json:"name"`
}

func (t *Todo) CreateTodo() {
  gdb.Create(&t)
}

func CreateTodo(name string) (*Todo){
  todo := &Todo{Name: name}
  todo.CreateTodo()
  return todo
}

func GetAllTodos() []Todo {
  todos := []Todo{}
  gdb.Find(&todos)
  return todos
}

func CreateTodoTable()  {
  gdb.CreateTable(&Todo{})
}
