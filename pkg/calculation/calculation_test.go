package calculation_test

import (
	"testing"

	"github.com/AlexDillz/Calc_server_yandex/pkg/calculation"
)

func TestCalc(t *testing.T) {
	successCases := []struct {
		name       string
		expression string
		expected   float64
	}{
		{"Simple addition", "1+1", 2},
		{"Multiplication and addition", "2+2*2", 6},
		{"Brackets with multiplication", "(2+2)*2", 8},
		{"Division", "8/2", 2},
	}

	for _, test := range successCases {
		t.Run(test.name, func(t *testing.T) {
			result, err := calculation.Calc(test.expression)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != test.expected {
				t.Errorf("expected %f, got %f", test.expected, result)
			}
		})
	}

	errorCases := []struct {
		name       string
		expression string
	}{
		{"Empty expression", ""},
		{"Invalid characters", "1+2a"},
		{"Unmatched brackets", "(1+2"},
		{"Division by zero", "1/0"},
	}

	for _, test := range errorCases {
		t.Run(test.name, func(t *testing.T) {
			_, err := calculation.Calc(test.expression)
			if err == nil {
				t.Fatalf("expected error for expression %s, got none", test.expression)
			}
		})
	}
}