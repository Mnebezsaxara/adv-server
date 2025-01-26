package main

import (
	"testing"
)

func TestUserModel(t *testing.T) {
	user := User{Email: "test@example.com", Password: "123456"}
	if user.Email != "test@example.com" || user.Password != "123456" {
		t.Errorf("User model fields do not match. Got: %v", user)
	}
}

func TestBookingModel(t *testing.T) {
	booking := Booking{Date: "2025-01-22", Time: "10:00", Field: "Football Field 1"}
	if booking.Date != "2025-01-22" || booking.Time != "10:00" || booking.Field != "Football Field 1" {
		t.Errorf("Booking model fields do not match. Got: %v", booking)
	}
}
