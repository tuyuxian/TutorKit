package main

import (
	"context"
	"log"
	"os"

	"backend/internal/ent"
	"backend/internal/ent/migrate"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Load env config
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	// Set up MySQL config
	mc := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "0.0.0.0" + ":" + os.Getenv("DB_PORT"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	client, err := ent.Open("mysql", mc.FormatDSN())
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Create DB schema with global ID
	if err := client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
