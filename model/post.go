package model

import "time"

type Post struct {
	Code        string `json:"code"`
	Title       string `json:"title"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Type        string `json:"type"`
	Description string `json:"description"`
	DateCreated time.Time 
	LastUpdated time.Time 
}
