package domain

type User struct{
	FirstName      string     `json:"first_name" gorm:"not null" binding:"required" form:"first_name"`
	LastName       string     `json:"last_name" gorm:"not null" binding:"required" form:"last_name"`
	Email          string     `json:"email" gorm:"not null" binding:"required" form:"email"`
	Password       string     `json:"password" gorm:"-" binding:"required" form:"password"`
	HashedPassword string     `json:"-,omitempty" gorm:"not null"`
	Phone          string     `json:"phone,omitempty"`
	ImageUrl       string     `json:"image_url,omitempty"`
}