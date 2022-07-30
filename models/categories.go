package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Categories struct {
	ID          string    `json:"id" gorm:"primary_key"`
	Name        string    `json:"name" gorm:"type:varchar(255)"`
	Description string    `json:"description" gorm:"type:text"`
	Created_at  time.Time `json:"created_at"`
	Created_by  string    `json:"created_by" gorm:"type:varchar(50)"`
	Updated_at  time.Time `json:"updated_at"`
	Updated_by  string    `json:"updated_by" gorm:"type:varchar(50)"`
	Deleted_at  time.Time `json:"deleted_at"`
	Deleted_by  string    `json:"deleted_by" gorm:"type:varchar(50)"`
}

func (m *Categories) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.Created_at = time.Now()
	return nil
}

func (m *Categories) BeforeUpdate(db *gorm.DB) error {
	m.Updated_at = time.Now()
	return nil
}

func (m *Categories) BeforeDelete(db *gorm.DB) error {
	m.Deleted_at = time.Now()
	return nil
}
