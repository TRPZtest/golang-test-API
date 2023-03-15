package dal

import (
	"log"
	models "testAPI/models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type UrlRepositoryInterface interface {
	GetUrlPackages(idList []uint) ([]models.UrlPackage, error)
}

type UrlRepository struct {
	db gorm.DB
}

type connectionFailed struct{}

func (connectionFailed) Error() string {
	return "Failed to connect to database"
}

func NewUrlRepository() (UrlRepository, error) {
	dsn := "Server=db-test;Database=testDb;User=sa;Password=S3cur3P@ssW0rd!;"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(connectionFailed{})
		return UrlRepository{db: *db}, connectionFailed{}
	}
	return UrlRepository{db: *db}, err
}

func (repository *UrlRepository) GetUrlPackages(idList []uint) ([]models.UrlPackage, error) {
	var urlPackages []models.UrlPackage

	tx := repository.db.Find(&urlPackages, idList)

	return urlPackages, tx.Error
}
