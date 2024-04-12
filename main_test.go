package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name     string
		val1     int
		val2     int
		expected int
	}{
		{
			name:     "success to sum values",
			val1:     5,
			val2:     5,
			expected: 10,
		},
		{
			name:     "fail to sum values",
			val1:     0,
			val2:     0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sum(tt.val1, tt.val2)
			assert.Equal(t, tt.expected, result)
		})
	}
}
