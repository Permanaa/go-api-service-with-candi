package domain

import "github.com/golangid/candi/candishared"

// FilterAuth model
type FilterAuth struct {
	candishared.Filter
	ID        *int     `json:"id"`
	StartDate string   `json:"startDate"`
	EndDate   string   `json:"endDate"`
	Preloads  []string `json:"-"`
	Email     string   `json:"email"`
}
