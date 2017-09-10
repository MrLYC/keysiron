package storage

import (
	"github.com/jinzhu/gorm"
	"time"
)

type BaseModel struct {
	gorm.Model

	CreateTime     time.Time
	LastModidyTime time.Time
}

type User struct {
	BaseModel

	ID       int    `gorm:"AUTO_INCREMENT"`
	Name     string `gorm:"size:64"`
	NickName string `gorm:"size:255"`
}

type Group struct {
	BaseModel

	ID             int    `gorm:"AUTO_INCREMENT"`
	Name           string `gorm:"size:64"`
	Desc           string `gorm:"size:255"`
	CreateTime     time.Time
	LastModidyTime time.Time
}

type UserGroup struct {
	BaseModel

	User  User  `gorm:"ForeignKey:user_id;AssociationForeignKey:UserRef"`
	Group Group `gorm:"ForeignKey:group_id;AssociationForeignKey:GroupRef"`
}
