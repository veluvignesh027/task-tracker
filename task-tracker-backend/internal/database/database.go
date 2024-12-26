package database

import (
	"fmt"
	"log"
	"os"
	"task-tracker/internal/models"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Service represents a service that interacts with a database.
type Service interface {
   GetDBInstance()  *gorm.DB 
}

type service struct {
	db *gorm.DB
}

var (
	database   = os.Getenv("BLUEPRINT_DB_DATABASE")
	password   = os.Getenv("BLUEPRINT_DB_PASSWORD")
	username   = os.Getenv("BLUEPRINT_DB_USERNAME")
	port       = os.Getenv("BLUEPRINT_DB_PORT")
	host       = os.Getenv("BLUEPRINT_DB_HOST")
	schema     = os.Getenv("BLUEPRINT_DB_SCHEMA")
	dbInstance *service
)

func New() Service {
	if dbInstance != nil {
		return dbInstance
	}

    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s", host, port, username, password, database, schema)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

    /*connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}*/
	dbInstance = &service{
		db: db,
	}
	
    err = dbInstance.CreateAllSchemas(&models.User{}, &models.Ticket{}, &models.Story{})
    if err != nil{
        log.Fatal(err)
    }

    return dbInstance
}

func (s *service) CreateAllSchemas(dst ...interface{})error{
    err := s.db.AutoMigrate(dst...)
    if err != nil{
        log.Println(err)
    }
    return err
}

func (s *service) GetDBInstance() *gorm.DB{
    return s.db
}
