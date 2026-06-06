package models

type EndpointDoc struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Category    string `json:"category"`
	Endpoint    string `json:"endpoint"`
	Method      string `json:"method"`
	Description string `json:"description"`
}