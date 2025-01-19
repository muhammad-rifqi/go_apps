package migration

import (
	"fmt"

	"github.com/muhammad-rifqi/go_apps/database"
	"github.com/muhammad-rifqi/go_apps/models/entity"
)

func RunMigration() {
	err := database.Database.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Create Table")
}
