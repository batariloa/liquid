package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

var (
	DB *sql.DB
)

func Init() error {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	err = createSchema(conn)
	if err != nil {
		log.Fatal("Error executing SQL file:", err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	log.Println("Successfully connected to the database!")

	DB = conn
	return nil
}

func createSchema(db *sql.DB) error {
	schemaSQL := `
		CREATE TABLE IF NOT EXISTS artists (
		    id SERIAL PRIMARY KEY,
		    NAME TEXT NOT NULL 
		);

		CREATE TABLE IF NOT EXISTS songs (
    	id SERIAL PRIMARY KEY,
   		file_path TEXT NOT NULL,
    	title TEXT NOT NULL,
    	artist INT NOT NULL, 
    	FOREIGN KEY (artist) REFERENCES artists(id));

`
	_, err := db.Exec(schemaSQL)
	if err != nil {
		return fmt.Errorf("apierror executing schema creation: %w", err)
	}

	return nil
}
