package data

import (
	"github.com/haappywater/virtualwallet/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Data struct {
	db *gorm.DB
}

func NewData(db *gorm.DB) *Data {
	return &Data{db: db}
}

func NewGormDB(conf *conf.Data) (*gorm.DB, func(), error) {

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		return nil, func() {}, err
	}

	db.AutoMigrate(
		&VirtualWallet{},
		&Transcation{},
	)

	return db, func() {
		// TODO: 关闭连接
	}, nil
}
