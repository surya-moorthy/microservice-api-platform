package database

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type databaseConfig struct {
	Host       string
	Port       int
	Username   string
	Password   string
	DBName     string
	SSLMode    string
	MaxOpen    int
	Maxidle    int
	IdleTimOut string
}

func loadDatabaseConfig() databaseConfig {
	viper.SetConfigName("database")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading a config file: %v", err)
	}
	return databaseConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DBName:   viper.GetString("database.dbname"),
		SSLMode:  viper.GetString("database.sslmode"),
		MaxOpen:  viper.GetInt("database.maxopen"),
		Maxidle:  viper.GetInt("database.maxidle"),
		IdleTimOut:  viper.GetString("database.idletimeOut"),
	}
}
func NewDBconnection() (*gorm.DB, error) {
	config := loadDatabaseConfig()

	dsn := fmt.Sprintf(
		"host=%s,\n port=%d,\n username=%s,\n password=%s\n dbname=%s\n sslmode=%s\n  maxopen=%d\n maxidle=%d\n idleTImeout=%s\n",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.DBName,
		config.SSLMode,
		config.MaxOpen,
		config.Maxidle,
		config.IdleTimOut,

	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database due to :%v", err)
	}
    psql,err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance due to :%v", err)
	}
    psql.SetConnMaxLifetime(time.Hour)
	return db, nil
}
