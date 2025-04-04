package model

import "time"

type Person struct {
	Id        int64     `gorm:"primary key; serial" json:"id"`
	Name      string    `gorm:"type: text;not null" json:"name"`
	Email     string    `gorm:"type: text;not null;uniq" json:"email"`
	Password  string    `gorm:"type: text;not null" json:"password"`
	IsActive  bool      `gorm:"type: boolean;default:true" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Person) TableName() string { return "persons" }
