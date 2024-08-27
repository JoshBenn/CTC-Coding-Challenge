package common

// API paths
type Path string

const (
	// User: Registration[PUT] and Login[POST]
	Authentication Path = "/login"

	// Message: New[POST] Refresh[GET]
	Chat Path = "/chat"
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
	MethodNotAllowed JsonComponent = "Method not allowed"
	Success          JsonComponent = "STATUS_SUCCESS"
	Fail             JsonComponent = "STATUS_FAILURE"
)

// Central source of truth for all cookie components
type CookieComponent string

const (
	AuthToken CookieComponent = "auth_token"
)
