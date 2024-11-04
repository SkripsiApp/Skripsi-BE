package model

type Admin struct {
	Id       string `gorm:"primaryKey;type:varchar(191);not null" json:"id"`
	Email    string `gorm:"unique;not null"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
