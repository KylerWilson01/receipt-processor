// Package models holds all the structs for this project
package models

// Item is what is used to make up a receipt
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}
