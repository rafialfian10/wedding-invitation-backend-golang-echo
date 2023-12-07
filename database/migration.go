package database

import (
	"fmt"
	"wedding/models"
	"wedding/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Pricing{},
		&models.Content{},
		&models.Feature{},
		&models.Header{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
