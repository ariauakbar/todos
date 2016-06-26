package models

import (
  // "database/sql"
  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
)

func GetDBConnection() (*gorm.DB, error) {
  db, err := gorm.Open("sqlite3", "../models/task.db")
  return db, err
}
