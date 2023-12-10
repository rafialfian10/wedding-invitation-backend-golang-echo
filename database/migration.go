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
		&models.Navigation{},
		&models.Fag{},
		&models.FagContent{},
		&models.Option{},
		&models.Footer{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
