package domain

import (
	"time"
)

type User struct {
	ID uint
	Name string `gorm:"type:VARCHAR(30);not null"`
	Email string `gorm:"type:VARCHAR(30);not null;index:idx_name,unique"`
	Password string `gorm:"type:VARCHAR(100);not null"`
	Role string `gorm:"type:VARCHAR(15);not null"`
	Status string `gorm:"type:ENUM('ACTIVE', 'INACTIVE');not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) GetUserID() uint {
    return u.ID
}
