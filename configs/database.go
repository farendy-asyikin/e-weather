package configs

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"main.go/models"
	"os"
	"strconv"
	"time"
)

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	logLevelStr := os.Getenv("LOG_LEVEL")
	if logLevelStr == "" {
		logLevelStr = strconv.Itoa(int(logger.Info))
	}
	logLevel, err := strconv.Atoi(logLevelStr)
	if err != nil {
		panic("Failed to load log level")
	}
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,               // Slow SQL threshold
			LogLevel:      logger.LogLevel(logLevel), // Log level
			Colorful:      true,                      // Disable color
		},
	)

	gormConfig := &gorm.Config{
		// enhance performance config
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 dbLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), gormConfig)
	if err != nil {
		panic("Failed to create a connection to database")
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Device{},
		&models.Sensor{},
		&models.SensorValue{},
	)
	if err != nil {
		panic("Failed to migrate database")
	}

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to close connection to database")
	}
	sqlDB.Close()
}
