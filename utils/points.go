// Package utils holds utility functions for the rest of the program
package utils

import (
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/KylerWilson01/receipt-processor.git/models"
)

// PointUtil holds all the utils that handles points
type PointUtil struct{}

// CheckRetailerName gets one point for every alphanumeric character in the retailer name.
func (PointUtil) CheckRetailerName(n string) int {
	points := 0

	for _, c := range n {
		if unicode.IsLetter(c) {
			points++
		} else if unicode.IsNumber(c) {
			points++
		}
	}

	return points
}

// CheckRoundDollar 50 points if the total is a round dollar amount with no cents.
func (PointUtil) CheckRoundDollar(a string) int {
	price, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0
	}

	// The total should never be negative as that would be a return and they shouldn't earn points on that
	if !(price > 0.0) {
		return 0
	}

	// If the price is a round dollar, this mod check should catch it
	//	variable can be inlined but for readability I chose not to
	isRoundDollar := math.Mod(price, 1.0) == 0
	if isRoundDollar {
		return 50
	}
	return 0
}

// CheckMultiple 25 points if the total is a multiple of 0.25.
func (PointUtil) CheckMultiple(a string) int {
	price, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0
	}

	// The total should never be negative as that would be a return and they shouldn't earn points on that
	if !(price > 0.0) {
		return 0
	}

	// If the price is a muliple of .25, this mod check should catch it
	//	variable can be inlined but for readability I chose not to
	isMultipleOfAQuarter := math.Mod(price, 0.25) == 0
	if isMultipleOfAQuarter {
		return 25
	}
	return 0
}

// CountLengthOfItems gets 5 points for every two items on the receipt.
func (PointUtil) CountLengthOfItems(i []models.Item) int {
	return int(math.Floor(float64(len(i)/2))) * 5
}

// CheckDescriptionLength If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
func (PointUtil) CheckDescriptionLength(i models.Item) int {
	multOfThree := (len(strings.TrimSpace(i.ShortDescription)) % 3) == 0

	if !multOfThree {
		return 0
	}

	p, err := strconv.ParseFloat(i.Price, 64)
	if err != nil {
		return 0
	}

	return int(math.Ceil(p * 0.2))
}

// CheckDate 6 points if the day in the purchase date is odd.
func (PointUtil) CheckDate(d string) int {
	const shortForm = "2006-01-02"
	date, err := time.Parse(shortForm, d)
	if err != nil {
		return 0
	}

	oddDay := date.Day()%2 == 1
	if oddDay {
		return 6
	}

	return 0
}

// CheckTime 10 points if the time of purchase is after 2:00pm and before 4:00pm.
func (PointUtil) CheckTime(t string) int {
	const shortForm = "15:04"
	date, err := time.Parse(shortForm, t)
	if err != nil {
		return 0
	}

	afterTwo := 14.0 < (float64(date.Hour()) + (float64(date.Minute()) / 60.0))
	beforeFour := (float64(date.Hour()) + (float64(date.Minute()) / 60.0)) < 16.0
	if afterTwo && beforeFour {
		return 10
	}

	return 0
}
