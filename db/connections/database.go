package database

import (
	"log"

	"github.com/ninsid711/CodeFusion-Judger/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "user=postgres.hvpgfjhgqzlpdsintxog password=Sidd7117*@@ host=aws-0-us-east-1.pooler.supabase.com port=6543 dbname=postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	db.AutoMigrate(&models.Ideas{})
	return db
}
