package storage

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	id   int
	name string
}
