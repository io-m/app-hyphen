package helpers

import (
	"crypto/ed25519"
	"errors"
	"fmt"
	"net/mail"
	"regexp"
)

var (
	hasLower    = regexp.MustCompile(`[a-z]`).MatchString
	hasUpper    = regexp.MustCompile(`[A-Z]`).MatchString
	hasNumber   = regexp.MustCompile(`[0-9]`).MatchString
	isValidName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateTextLength(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func ValidateName(value string) error {
	if err := ValidateTextLength(value, 2, 100); err != nil {
		return err
	}
	if !isValidName(value) {
		return fmt.Errorf("must contain only letters or spaces")
	}
	return nil
}

// Password must have at least 6 characters and at least one small letter, one capital letter and one number
func ValidatePassword(value string) error {
	if !hasLower(value) || !hasUpper(value) || !hasNumber(value) {
		return errors.New("password is not valid")
	}
	return ValidateTextLength(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateTextLength(value, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not a valid email address")
	}
	return nil
}

func IsPasetoKeyOfProperLength(private, public string) bool {
	if len(public) != ed25519.PublicKeySize*2 {
		return false
	}
	if len(private) != ed25519.PrivateKeySize*2 {
		return false
	}
	return true
}
