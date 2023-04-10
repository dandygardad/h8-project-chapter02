package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"project08/model/entity"
)

type Gorm struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string

	DB *gorm.DB
}

type GormDb struct {
	*Gorm
}

var (
	NewGorm *GormDb
)

func InitGorm() error {
	NewGorm = new(GormDb)
	NewGorm.Gorm = &Gorm{
		Username: os.Getenv("PGUSER"),
		Password: os.Getenv("PGPASSWORD"),
		Host:     os.Getenv("PGHOST"),
		DBName:   os.Getenv("PGDATABASE"),
		Port:     os.Getenv("PGPORT"),
	}

	err := NewGorm.Gorm.OpenConnection()
	if err != nil {
		return err
	}
	return nil
}

func (p *Gorm) OpenConnection() error {
	dbConfig := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", p.Host, p.Username, p.Password, p.DBName, p.Port)
	db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	if err != nil {
		return err
	}

	p.DB = db

	err = db.Debug().AutoMigrate(entity.Book{})
	if err != nil {
		return err
	}
	return nil
}
