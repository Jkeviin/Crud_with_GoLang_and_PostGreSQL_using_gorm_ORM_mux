package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirsName string `gorm:"not null"`
	LastName string `gorm:"not null"`
	Email    string `gorm:"not null;uniqueIndex"`
	Tasks    []Task // Uno a muchos, model Task
}
