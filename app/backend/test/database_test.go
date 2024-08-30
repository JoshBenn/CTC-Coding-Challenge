package test

import (
	"testing"
)

// This is intended to test the database, however, if one isn't set up on the system then it will not work
func TestDatabaseConnection(t *testing.T) {
	// // Set up environment variables
	// os.Setenv("LOG_FILE", "test.log")
	// defer os.Unsetenv("LOG_FILE")
	// os.Setenv("DATABASE_URI", "postgresql://postgres:db_ultramegasecretpassword@database:5432/postgres")
	// defer os.Unsetenv("DATABASE_URI")

	// node := common.NewNode()
	// defer node.Shutdown()
	// _, conn, err := database.NewDatabaseConnection(node)
	// if err != nil {
	// 	t.Error(err.Error())
	// }
	// conn.Close(context.Background())
}
