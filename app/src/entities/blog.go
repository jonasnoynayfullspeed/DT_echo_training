package entities

import (
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	ID          int64 `gorm:"primaryKey;autoIncrement"`
	Title       string
	Description string
	Content     string    `gorm:"type:LONGTEXT"`
	CreatedAt   time.Time //Populated by gorm.Model
	UpdatedAt   time.Time //Populated by gorm.Model
	DeletedAt   time.Time `gorm:"index;default:NULL"` //Populated by gorm.Model
}

type Blogs []Blog
