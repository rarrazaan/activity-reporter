package dependency

import (
	"fmt"
	"mini-socmed/internal/cons"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config Config, logger Logger) *gorm.DB {
	dsn := fmt.Sprintf(cons.ConnectionStringTemplate,
		config.PostgreDB.DBHost,
		config.PostgreDB.DBUser,
		config.PostgreDB.DBPass,
		config.PostgreDB.DBName,
		config.PostgreDB.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		logger.Errorf("Error connecting to Database", err)
		return nil
	}

	logger.Infof("Successfully Connect to Database", nil)
	return db
}
