package models

import (
        "gorm.io/gorm"
)

type User struct {
        gorm.Model
    Name     string   `gorm:"not null" json:"name"`
        Email    string   `gorm:"unique;not null" json:"email"`
        Password string   `gorm:"not null" json:"password"`
}

type Story struct {
        gorm.Model
        Name          string `gorm:"not null"`
        StoryID      int    `gorm:"unique;not null"`
        UserCreatedID uint  `gorm:"not null"`
        UserAssignedID uint  `gorm:"not null"`
    Description string
    Status string `gorm:"not null"`
    Priority string `gorm:"not null"`
}

type Ticket struct {
        gorm.Model
        Name          string `gorm:"not null"`
        TicketID      int    `gorm:"unique;not null"`
        UserCreatedID uint
        UserAssignedID uint
        Status        string `gorm:"not null"`
}

