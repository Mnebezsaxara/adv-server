package main

import (
	"testing"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func TestE2EBooking(t *testing.T) {
	caps := selenium.Capabilities{"browserName": "chrome"}
	chromeCaps := chrome.Capabilities{Args: []string{"--disable-gpu", "--no-sandbox"}}
	caps.AddChrome(chromeCaps)

	wd, err := selenium.NewRemote(caps, "http://localhost:4444")
	if err != nil {
		t.Fatalf("Failed to connect to WebDriver: %v", err)
	}
	defer wd.Quit()

	// Открытие страницы авторизации
	err = wd.Get("http://127.0.0.1:5500/SportLife-SportComplex/assignment1/front/form.html")
	if err != nil {
		t.Fatalf("Failed to load form.html: %v", err)
	}

	// Добавляем паузу для наглядности
	time.Sleep(2 * time.Second)

	// Тест авторизации
emailField, err := wd.FindElement(selenium.ByID, "email")
if err != nil {
    t.Fatalf("Failed to find email field: %v", err)
}
emailField.SendKeys("user1@example.com")

passwordField, err := wd.FindElement(selenium.ByID, "password")
if err != nil {
    t.Fatalf("Failed to find password field: %v", err)
}
passwordField.SendKeys("password123")

loginButton, err := wd.FindElement(selenium.ByCSSSelector, "#login-form button[type='submit']")
if err != nil {
    t.Fatalf("Failed to find login button: %v", err)
}
loginButton.Click()

// Подождем, чтобы alert мог появиться
time.Sleep(2 * time.Second)

// Обработка alert после авторизации
alertText, err := wd.AlertText()
if err == nil {
    if alertText != "Login successful" {
        t.Errorf("Unexpected alert message. Got: %s, expected: %s", alertText, "Login successful")
    }
    err = wd.AcceptAlert() // Закрытие alert
    if err != nil {
        t.Fatalf("Failed to accept alert: %v", err)
    }
} else {
    t.Fatalf("Failed to handle alert or no alert present: %v", err)
}


	// Открытие страницы бронирования
	err = wd.Get("http://127.0.0.1:5500/SportLife-SportComplex/assignment1/front/booking.html")
	if err != nil {
		t.Fatalf("Failed to load booking.html: %v", err)
	}
	time.Sleep(2 * time.Second)

	// Тест создания бронирования
	dateField, err := wd.FindElement(selenium.ByID, "date")
	if err != nil {
		t.Fatalf("Failed to find date field: %v", err)
	}
	dateField.SendKeys("2025-01-22")
	time.Sleep(1 * time.Second)

	timeField, err := wd.FindElement(selenium.ByID, "time")
	if err != nil {
		t.Fatalf("Failed to find time field: %v", err)
	}
	timeField.SendKeys("10:00")
	time.Sleep(1 * time.Second)

	fieldDropdown, err := wd.FindElement(selenium.ByID, "field")
	if err != nil {
		t.Fatalf("Failed to find field dropdown: %v", err)
	}
	fieldDropdown.SendKeys("Поле Бекет Батыра")
	time.Sleep(1 * time.Second)

	submitButton, err := wd.FindElement(selenium.ByCSSSelector, "#booking-form button[type='submit']")
	if err != nil {
		t.Fatalf("Failed to find submit button: %v", err)
	}
	submitButton.Click()
	time.Sleep(2 * time.Second)

	// Ожидание появления alert
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		alertText, err := wd.AlertText()
		if err == nil {
			if alertText != "Booking created" {
				t.Errorf("Unexpected alert message. Got: %s, expected: %s", alertText, "Booking created")
			}
			wd.AcceptAlert() // Закрытие alert
			return
		}
	}

	t.Fatalf("Failed to handle alert or no alert present after retries")
}
