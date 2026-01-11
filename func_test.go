package main

import (
	"testing"
)

func TestPaintNeeded(t *testing.T) {
	// Test valid inputs
	result, err := paintNeeded(10, 10)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expected := 10.0 // (10 * 10) / 10 = 10
	if result != expected {
		t.Errorf("got %.2f, want %.2f", result, expected)
	}
}

func TestPaintNeededNegativeWidth(t *testing.T) {
	_, err := paintNeeded(-5, 10)
	if err == nil {
		t.Error("expected error for negative width, got nil")
	}
}

func TestPaintNeededNegativeHeight(t *testing.T) {
	_, err := paintNeeded(10, -5)
	if err == nil {
		t.Error("expected error for negative height, got nil")
	}
}

// Table-driven tests (Go best practice)
func TestPaintNeededTableDriven(t *testing.T) {
	tests := []struct {
		name      string
		width     float64
		height    float64
		want      float64
		wantError bool
	}{
		{"valid dimensions", 4.0, 5.0, 2.0, false},
		{"zero dimensions", 0, 0, 0, false},
		{"negative width", -1, 5, 0, true},
		{"negative height", 5, -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := paintNeeded(tt.width, tt.height)
			if (err != nil) != tt.wantError {
				t.Errorf("error = %v, wantError = %v", err, tt.wantError)
				return
			}
			if got != tt.want {
				t.Errorf("got %.2f, want %.2f", got, tt.want)
			}
		})
	}
}
