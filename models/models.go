package models

import "gorm.io/gorm"

type Grade struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey" json:"id"`
	StudentID int     `gorm:"not null" json:"studentId"`
	Code      string  `gorm:"not null" json:"code"`
	Value     float64 `gorm:"not null" json:"value"`
}

type Student struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey" json:"id"`
	Name      string  `gorm:"not null" json:"name"`
	Surname   string  `gorm:"not null" json:"surname"`
	StdNumber string  `gorm:"not null unique" json:"stdNumber"`
	Grades    []Grade `gorm:"foreignKey:StudentID" json:"grades"`
}
