package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"number 1", 1, "1"},
		{"number 2", 2, "2"},
		{"Fizz for 3", 3, "Fizz"},
		{"Fizz for 6", 6, "Fizz"},
		{"Buzz for 5", 5, "Buzz"},
		{"Buzz for 10", 10, "Buzz"},
		{"FizzBuzz for 15", 15, "FizzBuzz"},
		{"FizzBuzz for 0", 0, "FizzBuzz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FizzBuzz(tt.input)
			if got != tt.expected {
				t.Errorf("FizzBuzz(%d) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
