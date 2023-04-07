package util

import "unicode"

// gotta rename it....
type IValidator interface {
	ValidatePassword(password string) bool
}

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) ValidatePassword(password string) bool {
	hasDigit, hasUpper, hasLower := false, false, false
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			hasDigit = true
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case c == ' ':
			return false
		}
	}
	return hasDigit && hasUpper && hasLower && len(password) >= 8 && len(password) <= 32
}
