package models

import "time"

type User struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (user *User) TableName() string {
	return "user"
}

type UserLogin struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UserRegister struct {
	Email     string `form:"email" json:"email" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	FirstName string `form:"first_name"`
	LastName  string `form:"last_name"`
}

func (user *User) ResponseMap() map[string]interface{} {
	res := make(map[string]interface{})
	res["id"] = user.ID
	res["email"] = user.Email
	res["first_name"] = user.FirstName
	res["last_name"] = user.LastName
	res["is_active"] = user.IsActive
	res["created_at"] = user.CreatedAt
	res["updated_at"] = user.UpdatedAt
	return res
}
