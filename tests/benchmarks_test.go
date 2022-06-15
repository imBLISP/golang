package tests

import (
	"testing"
)

const (
	Email  = "email"
	Push   = "push"
	SMS    = "sms"
	In_App = "in_app"
	Voice  = "voice"
)

func isFieldInValidArray(field string, validInput [5]string) bool {
	for _, v := range validInput {
		if field == v {
			return true
		}
	}

	return false
}

func isFieldInValidHash(field string, validInput map[string]bool) bool {
	if _, ok := validInput[field]; ok {
		return true
	}

	return false
}

func isFieldInValidSlice(field string, validInput []string) bool {
	for _, v := range validInput {
		if field == v {
			return true
		}
	}

	return false
}

var fields = [...]string{"email", "push", "sms", "in_app", "voice"}

func BenchmarkHash(b *testing.B) {
	validInput := map[string]bool{Email: true, Push: true, SMS: true, In_App: true, Voice: true}

	for _, f := range fields {
		isFieldInValidHash(f, validInput)
	}
}

func BenchmarkArray(b *testing.B) {
	validInput := [...]string{Email, Push, SMS, In_App, Voice}

	for _, f := range fields {
		isFieldInValidArray(f, validInput)
	}
}

func BenchmarkSlice(b *testing.B) {
	validInput := []string{Email, Push, SMS, In_App, Voice}

	for _, f := range fields {
		isFieldInValidSlice(f, validInput)
	}
}
