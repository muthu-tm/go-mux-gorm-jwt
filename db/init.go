package db

import (
	"log"

	"github.com/jinzhu/gorm"
	"go-rest-examples/models"
)

var users = []models.User{
	models.User{
		Name: "Muthu T",
		Email:    "example@gmail.com",
		Password: "password",
	},
}

func SeedData(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}