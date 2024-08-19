package common

// API paths
type Path string

const (
	// User: registration[PUT] and log in/out[POST]
	Authentication Path = "/login"

	// Message: New/Update[Post]
	Message Path = "/message"
)

// Central source of truth for all json components
type JsonComponent string

const (
	ContentType      JsonComponent = "Content-Type"
	ApplicationJson  JsonComponent = "application/json"
	Allow            JsonComponent = "Allow"
	Get              JsonComponent = "GET"
	Post             JsonComponent = "POST"
	Put              JsonComponent = "PUT"
	MethodNotAllowed JsonComponent = "Mmethod not allowed"
)
