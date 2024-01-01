package storage

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

type Storage struct {
    db *gorm.DB
}

func New(path string) (*Storage, error) {
    db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Set the maximum open connections (optional, but you can do it)
    sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }
    sqlDB.SetMaxOpenConns(1)

    // Auto Migrate your tables (similar to running migrations)
    err = db.AutoMigrate(&YourModel{}) // Replace YourModel with your actual model struct
    if err != nil {
        return nil, err
    }

    return &Storage{db: db}, nil
}
