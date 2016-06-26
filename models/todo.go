package models

import (
  "github.com/jinzhu/gorm"
)

var gdb, _ = GetDBConnection()

type Todo struct {
  gorm.Model
  Name string `sql:"type:varchar(200)" json:"name"`
  Finished bool `sql:"default:false" json:"finished"`
}

func (t *Todo) CreateTodo() {
  gdb.Create(&t)
}

func (t *Todo) ToggleFinished() {
  t.Finished = !t.Finished
  gdb.Save(&t)
}

func ToggleFinishedByID(id int) {
  t := &Todo{}
  gdb.Debug().Find( &t, id)
  t.ToggleFinished()
}

func CreateTodo(name string) (*Todo){
  todo := &Todo{Name: name}
  todo.CreateTodo()
  return todo
}

func GetAllTodos() []Todo {
  todos := []Todo{}
  gdb.Debug().Order("created_at desc").Find(&todos)
  return todos
}

func CreateTodoTable()  {
  gdb.CreateTable(&Todo{})
  gdb.AutoMigrate(&Todo{})
}
