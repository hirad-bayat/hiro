package Models

import "time"

type AccessToken struct {
	ID        uint      `gorm:"column:id; primaryKey"`
	Token     string    `gorm:"column:token; uniqueIndex; size:500"`
	UserID    uint      `gorm:"column:user_id"`
	ExpiresAt time.Time `gorm:"column:expires_at"`
	Revoked   bool      `gorm:"column:revoked; default:false"`
	User      User      `gorm:"foreignKey:UserID"`
}
