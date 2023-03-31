package config

import (
	"fmt"
	"os"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	User     string
	Password string
	Host     string
	Port     string
	DbName   string
}

func NewDatabase() (*Database, error) {
	var err error
	cfgDb := Database{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DbName:   os.Getenv("DB_NAME"),
	}
	err = cfgDb.validate()
	if err != nil {
		return nil, err
	}
	return &cfgDb, err
}

func (db Database) validate() error {
	return validation.ValidateStruct(
		&db,
		validation.Field(&db.User, validation.Required),
		validation.Field(&db.Password, validation.Required),
		validation.Field(&db.Host, validation.Required),
		validation.Field(&db.Port, validation.Required),
		validation.Field(&db.DbName, validation.Required),
	)
}

func MakeConnectionDatabase(data *Database, model ...interface{}) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s "+
		"password=%s "+
		"host=%s "+
		"TimeZone=Asia/Singapore "+
		"port=%s "+
		"dbname=%s", data.User, data.Password, data.Host, data.Port, data.DbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(model...); err != nil {
		return nil, err
	}
	return db, nil
}
