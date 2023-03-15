package main

import (
	"log"
	models "migrations/models"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {

	db := tryToConnect()
	tx := db.Exec("IF DB_ID('testDb') IS NULL CREATE DATABASE \"testDb\"")
	if tx.Error != nil {
		panic("failed to CREATE DATABASE testDb")
	}

	tx = db.Begin()

	tx = tx.Exec("USE testDb")
	if tx.Error != nil {
		panic("failed to use testDbest")
	}

	tx = tx.Exec("IF OBJECT_ID('dbo.UrlPackage', 'U') IS NOT NULL DROP TABLE dbo.UrlPackage;")

	err := tx.AutoMigrate(&models.UrlPackage{})
	if err != nil {
		panic("migration failure")
	} else {
		log.Println("migration success")
	}

	urlPackages := []models.UrlPackage{
		{Url: "http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=1"},
		{Url: "http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=2"},
		{Url: "http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=3"},
		{Url: "http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=4"},
		{Url: "http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=5"},
	}
	tx = tx.Create(urlPackages)
	if tx.Error != nil {
		panic("failed to add urlPackages")
	}
	tx.Commit()
}

func tryToConnect() *gorm.DB {
	isConnected := false

	for isConnected != true {
		dsn := "Server=db-test;Database=master;User=sa;Password=S3cur3P@ssW0rd!;"
		db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Println("failed to connect database")
			time.Sleep(2 * time.Second)
		} else {
			log.Println("Connected to database")
			return db
		}
	}
	return nil
}
