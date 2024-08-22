package common

// For environment variables
type Env string

const (
	// Secret for the JWT
	JwtSecret Env = "JWT_SECRET"
	// URI for the database
	DatabaseUri Env = "DATABASE_URI"
	// Password for the database
	PostgresPassword Env = "POSTGRES_PASSWORD"
	// Address/port used for the database
	BackendAddress Env = "BACKEND_ADDRESS"
	// Gets the logfile name
	logFile Env = "LOG_FILE"
	// file permissions
	permissions = 0644
)

// Logging level
type Level uint8

const (
	Debug Level = 0
	Info  Level = 1
	Warn  Level = 2
	Error Level = 3
)

// Contains log information
type Log struct {
	Level   Level
	Message string
}

// Generates a new log for the logging output
func NewLog(level Level, message string) Log {
	return Log{Level: level, Message: message}
}
