package model

import "time"

type Company struct {
	Code        string
	Name        string
	City        string
	Country     string
	Address     string
	Description string
	DateCreated time.Time
	LastUpdated time.Time
}
