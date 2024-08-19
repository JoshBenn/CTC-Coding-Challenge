package services

import (
	"context"
	"fmt"

	"github.com/JoshBenn/CTC-Coding-Challenge/common"
	"github.com/JoshBenn/CTC-Coding-Challenge/database"
	"github.com/jackc/pgx/v5"
)

// Represents the database connection and queries provider
type Provider struct {
	Queries *database.Queries
	Conn    *pgx.Conn
}

// Creates the struct for easy user management -- Ideally there's a limited pool of these connections
func NewProvider(node *common.Node) (Provider, error) {
	queries, conn, err := database.NewDatabaseConnection(node)
	if err != nil {
		return Provider{}, err
	}

	return Provider{Queries: queries, Conn: conn}, nil
}

// Closes the connection to the database
func (dbConn *Provider) CloseDbConn(node *common.Node) error {
	// If the connection is already closed
	if dbConn == nil {
		return nil
	}

	// Attempt to close the connection
	if err := dbConn.Conn.Close(context.Background()); err != nil {
		node.Log <- common.NewLog(common.Error, fmt.Sprintf("Could not close connection to database: %v", err))
		node.Output <- fmt.Sprintf("Failed to close connection to database: %v", err)
		return err
	}

	// Connection closed
	node.Output <- "Database connection successfully closed\n"
	return nil
}
