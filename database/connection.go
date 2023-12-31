package database

import (
	"os"

	helper "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/helpers"
	model "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	databaseURI := make(chan string, 1)

	if os.Getenv("GO_ENV") != "production" {
		databaseURI <- helper.GodotEnv("DATABASE_URI_DEV")
	} else {
		databaseURI <- os.Getenv("DATABASE_URI_PROD")
	}

	db, err := gorm.Open(postgres.Open(<-databaseURI), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	}

	if os.Getenv("GO_ENV") != "production" {
		logrus.Info("Connection to Database Successfully")
	}

	err = db.AutoMigrate(
		&model.EntityUsers{},
		&model.EntityPhotos{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
