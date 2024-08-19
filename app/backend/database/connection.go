package database

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/JoshBenn/CTC-Coding-Challenge/common"
	"github.com/jackc/pgx/v5"
)

const Schema = "./schema.sql"

// For migrating the database if necessary
func migrate(migrationFile string, queries *Queries, context *context.Context) error {
	fileContents, err := os.ReadFile(migrationFile)
	if err != nil {
		return err
	}

	_, err = queries.db.Exec(*context, string(fileContents))
	if err != nil {
		return err
	}

	return nil
}

// Get a new database connection
func NewDatabaseConnection(node *common.Node) (*Queries, *pgx.Conn, error) {
	context := context.Background()
	connString := os.Getenv(string(common.DatabaseUri))
	if len(connString) == 0 {
		return nil, nil, errors.New("no database uri found in environment variables")
	}

	var (
		conn *pgx.Conn
		err  error
	)

	// Attempt 5 times
	for i := 1; i <= 5; i++ {
		node.Output <- fmt.Sprintf("[TRY %d] Attempting to connecting to %v", i, connString)
		conn, err = pgx.Connect(context, connString)

		// If successfully connected
		if err == nil {
			node.Output <- "Successfully connected to database"
			break
		}

		// Break out if the 5th failure
		if i == 5 {
			node.Log <- common.NewLog(common.Error, fmt.Sprintf("Could not connect to database: %v", err))
			return nil, nil, err
		}
		node.Output <- fmt.Sprintf("Failed to connect to database: %v\nAttempting again in 10 seconds", err)
		node.Log <- common.NewLog(common.Warn, fmt.Sprintf("Failed to connect to database, attempt %d", i))
		// Wait 5 seconds
		time.Sleep(time.Second * 5)
	}

	queries := New(conn)

	err = migrate(string(Schema), queries, &context)
	if err != nil {
		node.Log <- common.NewLog(common.Error, fmt.Sprintf("Failed to run migrations: %v", err))
		return nil, nil, err
	}

	return queries, conn, nil
}
