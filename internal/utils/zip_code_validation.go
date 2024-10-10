package utils

import (
	"errors"
	"regexp"
)

func IsValid(zipCode string) error {
	if len(zipCode) != 8 {
		return errors.New("invalid zipcode: must be 8 digits")
	}
	if matched, _ := regexp.MatchString("^[0-9]+$", zipCode); !matched {
		return errors.New("invalid zipcode: must contain only digits")
	}
	return nil
}
