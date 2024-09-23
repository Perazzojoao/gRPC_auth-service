package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()"`
	Email     string         `bson:"email" gorm:"unique;not null"`
	Password  string         `bson:"password" gorm:"not null"`
	Active    bool           `bson:"active" gorm:"default:false"`
	CreatedAt time.Time      `bson:"created_at"`
	UpdatedAt time.Time      `bson:"updated_at"`
	DeletedAt gorm.DeletedAt `bson:"deleted_at" gorm:"index"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.Id = uuid.New()
	return
}
