package models

type GenericError struct {
	Error string `json:"error"`
}

type StatusServiceUnavailableError struct {
	Error          string `json:"error"`
	RateLimitReset string `json:"rateLimitReset"`
}
