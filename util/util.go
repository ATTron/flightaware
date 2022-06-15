package util

import (
	"errors"
	"log"
	"strings"
)

// EndPoint -- the default API endpoint for flightaware
const EndPoint = "https://aeroapi.flightaware.com/aeroapi/"

// ErrNotFound -- the basic error when we fail to return data
var ErrNotFound = errors.New("cannot fetch at this time. please try again")

// Check -- used to check any error values
func Check(err error) {
	if err != nil {
		log.Fatal("Unable to complete request due to error! ", err)
	}
}

// Join - combine strings
func Join(strs ...string) string {
	var sb strings.Builder
	for _, q := range strs {
		sb.WriteString(q)
	}
	return sb.String()

}
