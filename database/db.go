package database

import (
    "desa-digital/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("desa_digital.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }

    // Migrasi
    DB.AutoMigrate(&models.User{}, &models.Course{})
    
    // Seed data awal
    seedData()
}

func seedData() {
    // Tambah kursus contoh
    courses := []models.Course{
        {Title: "Pengenalan Desa Digital", Description: "Belajar dasar teknologi desa.", Content: "Konten kursus 1..."},
        {Title: "Manajemen Pertanian Modern", Description: "Teknik pertanian digital.", Content: "Konten kursus 2..."},
    }
    for _, course := range courses {
        DB.FirstOrCreate(&course, models.Course{Title: course.Title})
    }
}
