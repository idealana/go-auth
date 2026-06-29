package domain

import (
	"time"
)

type RefreshToken struct {
	ID uint
	UserID uint `gorm:"not null;index:idx_user_id"`
	TokenHash string `gorm:"type:CHAR(64);not null;index:idx_token_hash,unique"`
	ExpiredAt time.Time `gorm:"not null;index:idx_expired_at"`
	RevokeAt *time.Time `gorm:"index:idx_revoke_at"`
	LastUsedAt *time.Time
	IPAddress *string `gorm:"type:VARCHAR(45);null"`
	UserAgent *string `gorm:"type:VARCHAR(255);null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
