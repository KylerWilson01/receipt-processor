package models

import (
	"time"

	"github.com/google/uuid"
)

// Receipt is what gets parsed from the user
type Receipt struct {
	ID           *uuid.UUID `json:"id"`
	Retailer     string     `json:"retailer"`
	PurchaseDate time.Time  `json:"purchaseDate"`
	PurchaseTime time.Time  `json:"purchaseTime"`
	Items        []Item     `json:"items"`
	Total        string     `json:"total"`
}
