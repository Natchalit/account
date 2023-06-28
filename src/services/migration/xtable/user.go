package xtable

import (
	"github.com/gofrs/uuid"
)

func (Users) TableName() string {
	return `users`
}

type Users struct {
	User_id  *uuid.UUID `gorm:"type:uuid;primary_key;" json:"user_id"`
	Username string     `gorm:"" json:"username"`
	Password string     `gorm:"" json:"password"`
}
