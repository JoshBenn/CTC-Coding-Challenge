package common

type Path string
type Env string
type JsonComponent string

const (
	// User registration[PUT] and log in/out[POST]
	Authentication Path = "/login"
	Message        Path = "/message"

	// Password for the JWT
	JwtPassword Env = "JWT_PASSWORD"
	// URI for the database
	DatabaseUri Env = "DATABASE_URI"
	// Password for the database
	DatabasePassword Env = "DATABASE_PASSWORD"

	ContentType      JsonComponent = "Content-Type"
	ApplicationJson  JsonComponent = "application/json"
	Allow            JsonComponent = "Allow"
	Get              JsonComponent = "GET"
	Post             JsonComponent = "POST"
	Put              JsonComponent = "PUT"
	MethodNotAllowed JsonComponent = "Mmethod not allowed"
)

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

func NewLog(level Level, message string) Log {
	return Log{Level: level, Message: message}
}
