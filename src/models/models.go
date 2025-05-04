package models

import (
	"github.com/ahsankoushik/mosquser/src/config"
	"github.com/ahsankoushik/mosquser/src/utils"
)

func Migrate() {
	utils.Logger("Starting migration...")
	db := config.ConnectDB()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Acl{})

	utils.Logger("Done migration")
}
