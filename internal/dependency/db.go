package dependency

import (
	"fmt"
	"mini-socmed/internal/constant"
	"mini-socmed/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config Config, logger Logger) *gorm.DB {
	dsn := fmt.Sprintf(constant.ConnectionStringTemplate,
		config.PostgreDB.DBHost,
		config.PostgreDB.DBUser,
		config.PostgreDB.DBPass,
		config.PostgreDB.DBName,
		config.PostgreDB.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		return nil
	}
	// // UNCOMMENT WHEN TRYING TO RESET DB
	// Down(db)

	db.AutoMigrate(model.User{}, model.Photo{}, model.Activity{}, model.UserPhoto{})

	return db
}

func Down(db *gorm.DB) {
	db.Migrator().DropTable(model.Activity{})
	db.Migrator().DropTable(model.Photo{})
	db.Migrator().DropTable(model.User{})
}