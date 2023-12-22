package models

import "time"

type Post struct {
	ID        int64     `gorm:"primary_key:auto_increment" json:"id"`
	Title     string    `gorm:"size:200" json:"title"`
	Body      string    `gorm:"size:3000" json:"body"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (post *Post) TableName() string {
	return "post"
}

func (post *Post) ResponseMap() map[string]interface{} {
	res := make(map[string]interface{})
	res["id"] = post.ID
	res["title"] = post.Title
	res["body"] = post.Body
	res["created_at"] = post.CreatedAt
	res["updated_at"] = post.UpdatedAt
	return res
}
