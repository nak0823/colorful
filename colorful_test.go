package colorful

import (
	"testing"
)

func TestHexToRGB(t *testing.T) {
	tests := []struct {
		color  Color
		expect RGB
	}{
		{AliceBlue, RGB{240, 248, 255}},
		{YellowGreen, RGB{154, 205, 50}},
	}

	for _, tc := range tests {
		result := tc.color.ToRGB()

		if result != tc.expect {
			t.Errorf("Test failed: expected RGB %v, got %v", tc.expect, result)
		}
	}
}

func TestColorFromHTML(t *testing.T) {
	tests := []struct {
		input    string
		expected Color
	}{
		{"#F0F8FF", Color(0xF0F8FF)}, // AliceBlue
		{"#9ACD32", Color(0x9ACD32)}, // YellowGreen
	}

	for _, test := range tests {
		result := ColorFromHTML(test.input)
		if result != test.expected {
			t.Errorf("ColorFromHTML(%s) returned %d, expected %d", test.input, result, test.expected)
		}
	}
}

func TestColorFromRGB(t *testing.T) {
	tests := []struct {
		r, g, b  int
		expected Color
	}{
		{240, 248, 255, Color(0xF0F8FF)}, // AliceBlue
		{154, 205, 50, Color(0x9ACD32)},  // YellowGreen
	}

	for _, test := range tests {
		result := ColorFromRGB(test.r, test.g, test.b)

		if result != test.expected {
			t.Errorf("ColorFromRGB(%d, %d, %d) returned %d, expected %d", test.r, test.g, test.b, result, test.expected)
		}
	}
}
