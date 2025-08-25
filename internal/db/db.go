package db

import (
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/muhammad21236/Go-gRPC-Service/internal/rocket"
	_ "github.com/lib/pq"
	"log"
)

type Store struct {
	db *sqlx.DB
}

// New returns a new store or error
func New() (Store, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	connectionString := "postgres://" + dbUsername + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbTable + "?sslmode=" + dbSSLMode
	
	// Add retry mechanism for database connection
	var db *sqlx.DB
	var err error
	maxRetries := 10
	retryDelay := 3 * time.Second

	for i := 0; i < maxRetries; i++ {
		db, err = sqlx.Connect("postgres", connectionString)
		if err != nil {
			log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
			time.Sleep(retryDelay)
			continue
		}
		
		// Test the connection
		err = db.Ping()
		if err == nil {
			break
		}
		log.Printf("Database not ready (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(retryDelay)
	}

	if err != nil {
		return Store{}, err
	}
	return Store{db: db}, nil
}

func (s Store) GetRocketByID(id string) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

func (s Store) DeleteRocket(id string) error {
	return nil
}