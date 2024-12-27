package models

import (
	"gorm.io/gorm"
)

type User struct {
    UserID   uint   `gorm:"primaryKey;autoIncrement;unique"`
    Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
    gorm.Model
}

type Story struct {
	StoryID        int    `gorm:"primaryKey;autoIncrement;unique"`
    Name           string `gorm:"not null"`
	UserCreatedID  uint   `gorm:"not null"`
	UserAssignedID uint   `gorm:"not null"`
	Description    string
	Status         string `gorm:"not null"`
	Priority       string `gorm:"not null"`
    gorm.Model
}

type Ticket struct {
	TicketID       int    `gorm:"primaryKey;autoIncrement;unique"`
    Name           string `gorm:"not null"`
	UserCreatedID  uint
	UserAssignedID uint
	Status         string `gorm:"not null"`
    gorm.Model
}
