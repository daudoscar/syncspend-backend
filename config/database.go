package config

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the database and tests blob storage connection
func ConnectDatabase() {
	if ENV == nil {
		LoadConfig()
	}

	// Connect to MySQL Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		ENV.DBUser, ENV.DBPassword, ENV.DBHost, ENV.DBPort, ENV.DBName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Println("Database connection successful")

	err = TestBlobStorageConnection()
	if err != nil {
		log.Fatalf("Failed to connect to Azure Blob Storage: %v", err)
	}
	log.Println("Azure Blob Storage connection successful")
}

func TestBlobStorageConnection() error {
	client := BlobClient
	containerName := ENV.AzureContainerName

	ctx := context.Background()
	pager := client.NewListBlobsFlatPager(containerName, nil)

	if pager.More() {
		_, err := pager.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("failed to access container or list blobs: %v", err)
		}
	}

	log.Println("Azure Blob Storage connection successful.")
	return nil
}
