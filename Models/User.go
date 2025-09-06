package Models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `json:"password"`
	Blogs    []Blog `json:"blogs"` // A user can have many blogs
}
