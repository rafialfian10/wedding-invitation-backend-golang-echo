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
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
