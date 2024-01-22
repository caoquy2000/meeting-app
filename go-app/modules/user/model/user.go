package usermodel

import (
	"time"

	"github.com/caoquy2000/meeting-app/common"
)

const EntityName = "User"

type User struct {
	common.SQLModel    `json:",inline"`
	Email              string     `json:"email" gorm:"column:email;"`
	Username           string     `json:"username" gorm:"column:username;"`
	password           string     `json:"-" gorm:"column:password;"`
	Method             string     `json:"method" gorm:"column:method;"`
	Discriminator      string     `json:"discriminator" gorm:"column:discriminator;"`
	GlobalName         string     `json:"global_name" gorm:"column:global_name;"`
	Avatar             string     `json:"avatar" gorm:"column:avatar;"`
	Banner             string     `json:"banner" gorm:"column:banner;"`
	System             bool       `json:"system" gorm:"column:system;"`
	MfaEnabled         bool       `json:"mfa_enabled" gorm:"column:mfa_enabeld;"`
	AccentColor        int8       `json:"accent_color" gorm:"column:accent_color;"`
	Locale             string     `json:"locale" gorm:"column:locale;"`
	Flags              int8       `json:"flags" gorm:"column:flags;"`
	PremiumType        int8       `json:"premium_type" gorm:"column:premium_type;"`
	PublicFlags        int8       `json:"public_flags" gorm:"column:public_flags;"`
	AvatarDecoration   string     `json:"avatar_decoration" gorm:"column:avatar_decoration;"`
	ConfirmationToken  string     `json:"confirmation_token" gorm:"column:confirmation_token;"`
	RqChangePasswordAt *time.Time `json:"rq_change_password_at" gorm:"column:rq_change_password_at;"`
}
