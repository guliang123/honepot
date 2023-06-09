package models

import (
	"decept-defense/pkg/configs"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Model struct {
	Id         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

// SetUp initializes the database instance
func SetUp() {

	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configs.GetSetting().Database.DBUser,
		configs.GetSetting().Database.DBPassword,
		configs.GetSetting().Database.DBHost+":"+configs.GetSetting().Database.DBPort,
		configs.GetSetting().Database.DBName)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		zap.L().Panic(err.Error())
	}
	if db == nil {
		zap.L().Panic("db nil")
	}

	if err = db.AutoMigrate(&User{}); err != nil {
		zap.L().Panic("AutoMigrate error:" + err.Error())
	}
	_ = (&User{}).CreateDefaultUser()

	if err = db.AutoMigrate(&Agent{}); err != nil {
		zap.L().Panic("AutoMigrate error:" + err.Error())
	}
	if err = db.AutoMigrate(&VirusRecord{}); err != nil {
		zap.L().Panic("AutoMigrate error:" + err.Error())
	}
	if err = db.AutoMigrate(&Honeypot{}); err != nil {
		zap.L().Panic("AutoMigrate error:" + err.Error())
	}
	if err = db.AutoMigrate(&Protocols{}); err != nil {
		zap.L().Panic("AutoMigrate error:" + err.Error())
	}
	_ = (&Protocols{}).CreateDefaultProtocol()

	if err = db.AutoMigrate(&ProtocolProxy{}); err != nil {
		zap.L().Panic("AutoMigrate error:" + err.Error())
	}
	if err = db.AutoMigrate(&TransparentProxy{}); err != nil {
		zap.L().Panic("AutoMigrate error:" + err.Error())
	}
	if err = db.AutoMigrate(&Images{}); err != nil {
		zap.L().Error("AutoMigrate error:" + err.Error())
	}
	_ = (&Images{}).CreateDefaultImage()
	if err = db.AutoMigrate(&Bait{}); err != nil {
		zap.L().Panic("AutoMigrate error:" + err.Error())
	}
	if err = db.AutoMigrate(&BaitTask{}); err != nil {
		zap.L().Panic("AutoMigrate error:" + err.Error())
	}
	if err = db.AutoMigrate(&ProtocolEvent{}); err != nil {
		zap.L().Panic("AutoMigrate error:" + err.Error())
	}
	if err = db.AutoMigrate(&TransparentEvent{}); err != nil {
		zap.L().Panic("AutoMigrate error:" + err.Error())
	}
	if err = db.AutoMigrate(&FalcoAttackEvent{}); err != nil {
		zap.L().Panic("AutoMigrate error:" + err.Error())
	}
	if err = db.AutoMigrate(&TokenTraceLog{}); err != nil {
		zap.L().Panic("AutoMigrate error:" + err.Error())
	}
	if err = db.AutoMigrate(&Setting{}); err != nil {
		zap.L().Panic("AutoMigrate error:" + err.Error())
	}
	if err = db.AutoMigrate(&CounterEvent{}); err != nil {
		zap.L().Panic("AutoMigrate error:" + err.Error())
	}
	(&Setting{}).CreateDefaultSetting()
}
