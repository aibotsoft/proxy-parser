package models

import (
	"time"
)

// Proxy from https://www.sslproxies.org/
type Proxy struct {
	Ip        string
	Port      string
	Code      string
	Country   string
	Anonymity string
	CreatedAt time.Time
	UpdatedAt time.Time
}
