package database

import (
    "gorm.io/gorm"
    "microservices-api-platform/internal/models"
)

func RunMigrations(db *gorm.DB) error {
    // Auto-migrate creates/updates tables based on defined models
    return db.AutoMigrate(
        &models.User{},
        &models.Product{},
        &models.Order{},
        // Add more models as needed
    )
}