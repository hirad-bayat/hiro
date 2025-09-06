package Models

type Blog struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Title   string `gorm:"column:title" json:"title"`
	Content string `gorm:"column:content" json:"content"`

	UserID uint `json:"user_id"` // Foreign key
	User   User `json:"user"`    // Belongs to one user
}
