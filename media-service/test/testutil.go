package test

import (
	"StorageService/internal/db"
	"context"
	"database/sql"
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func setUpDbContainer(t *testing.T) (*sql.DB, error) {
	_, currentFile, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(filepath.Dir(currentFile))
	schemaFilePath := filepath.Join(projectRoot, "sql", "schema.sql")

	ctx := context.Background()

	fmt.Println("Absolute path to schema file:", schemaFilePath)
	c, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithInitScripts(schemaFilePath),
		postgres.WithDatabase("users-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		panic(err)
	}
	connStr, err := c.ConnectionString(ctx, "sslmode=disable", "application_name=test")
	assert.NoError(t, err)

	// Connect to the PostgreSQL database
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Error connecting to the database: %v", err)
	}

	// Point global db variable to test container connection
	db.DB = conn

	// Return the connection so that it can be closed
	return conn, nil
}

func setUpRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}
