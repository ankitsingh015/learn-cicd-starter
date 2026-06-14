package auth

import (
    "errors"
    "net/http"
    "testing"
)

func TestGetAPIKey(t *testing.T) {
    // Sahi code (t.Fatal hata diya hai)

    // Case 1: Success Case (Sahi header)
    headers := http.Header{}
    headers.Set("Authorization", "ApiKey ankit-secret-123")
    key, err := GetAPIKey(headers)
    if err != nil {
        t.Fatalf("Sahi header par error nahi aana chahiye tha, par aaya: %v", err)
    }
    if key != "ankit-secret-123" {
        t.Errorf("Expected 'ankit-secret-123', got: %s", key)
    }

    // Case 2: Failure Case 1 (Header gayab hai)
    emptyHeaders := http.Header{}
    _, err = GetAPIKey(emptyHeaders)
    if !errors.Is(err, ErrNoAuthHeaderIncluded) {
        t.Errorf("Expected ErrNoAuthHeaderIncluded, got: %v", err)
    }

    // Case 3: Failure Case 2 (Header malformed/galat format mein hai)
    badHeaders := http.Header{}
    badHeaders.Set("Authorization", "Bearer wrong-format-key")
    _, err = GetAPIKey(badHeaders)
    if err == nil {
        t.Errorf("Galat format par error aana chahiye tha, par nahi aaya")
    }
}
