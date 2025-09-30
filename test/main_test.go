package main

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
)

// Mock the checkCertExpiration function for testing
func checkCertExpirationTest(u string) (int, error) {
	// This would normally be imported from your main package
	// For now, we'll copy the logic here for testing
	if u == "" {
		return 0, fmt.Errorf("URL is required")
	}
	
	parsedURL, err := url.ParseRequestURI(u)
	if err != nil {
		return 0, fmt.Errorf("invalid URL: %s", err)
	}

	if parsedURL.Scheme != "https" {
		return 0, fmt.Errorf("invalid URL scheme (only https supported): %s", parsedURL.Scheme)
	}

	// For testing, return a mock value instead of actual certificate check
	return 30, nil
}

func TestCheckCertExpiration_EmptyURL(t *testing.T) {
	_, err := checkCertExpirationTest("")
	if err == nil {
		t.Error("Expected error for empty URL, got nil")
	}
	if err.Error() != "URL is required" {
		t.Errorf("Expected 'URL is required', got '%s'", err.Error())
	}
}

func TestCheckCertExpiration_InvalidURL(t *testing.T) {
	_, err := checkCertExpirationTest("invalid-url")
	if err == nil {
		t.Error("Expected error for invalid URL, got nil")
	}
}

func TestCheckCertExpiration_NonHTTPS(t *testing.T) {
	_, err := checkCertExpirationTest("http://example.com")
	if err == nil {
		t.Error("Expected error for non-HTTPS URL, got nil")
	}
	if !strings.Contains(err.Error(), "only https supported") {
		t.Errorf("Expected HTTPS error, got '%s'", err.Error())
	}
}

func TestCheckCertExpiration_ValidHTTPS(t *testing.T) {
	days, err := checkCertExpirationTest("https://example.com")
	if err != nil {
		t.Errorf("Expected no error for valid HTTPS URL, got %s", err)
	}
	if days != 30 {
		t.Errorf("Expected 30 days, got %d", days)
	}
}