package storage

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"keysiron/config"
)

func connectMySQL() (db *gorm.DB, err error) {
	var connStr = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.KeysironConfigVar.Database.User, config.KeysironConfigVar.Database.Password, config.KeysironConfigVar.Database.Host, config.KeysironConfigVar.Database.Port, config.KeysironConfigVar.Database.Name)
	return gorm.Open("mysql", connStr)
}

// Connect : make a connection
func Connect() (db *gorm.DB, err error) {
	return connectMySQL()
}
