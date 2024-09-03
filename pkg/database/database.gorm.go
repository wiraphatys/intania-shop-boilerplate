package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/wiraphatys/intania-shop-boilerplate/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormDatabase(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.GetDb().Host,
		cfg.GetDb().User,
		cfg.GetDb().Password,
		cfg.GetDb().Name,
		cfg.GetDb().Port,
		cfg.GetDb().SSLMode,
		cfg.GetDb().Timezone,
	)

	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // Slow SQL threshold
			LogLevel:      logger.Error, // Log level
			Colorful:      true,         // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                                   logger,
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
