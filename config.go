package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDB() (*gorm.DB, error) {
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=your_postgres_username dbname=ruud_pen password=your_postgres_password sslmode=disable")
    if err != nil {
        return nil, err
    }
    return db, nil
}
