package models

import "time"

type EndpointDoc struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Category    string `json:"category"`
	Endpoint    string `json:"endpoint"`
	Method      string `json:"method"`
	Description string `json:"description"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Source  string      `json:"source"`
	Updated time.Time   `json:"updated"`
	Data    interface{} `json:"data"`
}
