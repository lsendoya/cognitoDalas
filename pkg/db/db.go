package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	secret "github.com/lsendoya/cognitoDalas/pkg/aws"
	"github.com/lsendoya/cognitoDalas/pkg/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectDB(cfg aws.Config) (*gorm.DB, error) {
	dbCredentials := secret.GetSecretManager(cfg, os.Getenv("SECRET_NAME"))

	var rdsJSON secret.SecretRDS
	if err := helper.UnmarshalJSON([]byte(dbCredentials), &rdsJSON); err != nil {
		return nil, fmt.Errorf("error reading secret credentials AWS RDS %w", err)
	}

	db, errConn := gorm.Open(postgres.Open(makeDSN(rdsJSON)), &gorm.Config{})
	if errConn != nil {

		return nil, fmt.Errorf("gorm.Open(), Failed to connect to database, %w", errConn)
	}

	log.Println("Connection opened to database")

	return db, nil
}

func makeDSN(s secret.SecretRDS) string {
	dbName := os.Getenv("DB_NAME")
	port := 5432
	host := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		host, port, s.Username, s.Password, dbName)

	return dsn
}
