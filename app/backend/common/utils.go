package common

// For environment variables
type Env string

const (
	// Password for the JWT
	JwtPassword Env = "JWT_PASSWORD"
	// URI for the database
	DatabaseUri Env = "DATABASE_URI"
	// Password for the database
	DatabasePassword Env = "DATABASE_PASSWORD"
)

// Logging level
type Level uint8

const (
	Debug Level = 0
	Info  Level = 1
	Warn  Level = 2
	Error Level = 3
)

type Log struct {
	Level   Level
	Message string
}

// Generates a new log for the logging output
func NewLog(level Level, message string) Log {
	return Log{Level: level, Message: message}
}
