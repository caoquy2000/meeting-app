package usermodel

import "github.com/caoquy2000/meeting-app/common"

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
}
