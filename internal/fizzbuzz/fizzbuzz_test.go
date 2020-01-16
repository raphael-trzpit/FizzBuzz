package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name         string
		fizzMultiple int
		buzzMultiple int
		limit        int
		fizzStr      string
		buzzStr      string
		expected     []string
	}{
		{
			"From 1 to 20, 3-Fizz, 5-Buzz",
			3, 5, 20, "Fizz", "Buzz",
			[]string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz"},
		},
		{
			"From 1 to 10, 2-Fizz, 4-Buzz",
			2, 4, 10, "Fizz", "Buzz",
			[]string{"1", "Fizz", "3", "FizzBuzz", "5", "Fizz", "7", "FizzBuzz", "9", "Fizz"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := FizzBuzz(test.fizzMultiple, test.buzzMultiple, test.limit, test.fizzStr, test.buzzStr)

			if len(output) != len(test.expected) {
				t.Fatalf("Generated FizzBuzz list doesn't have the expected length. (expected: %d, got: %d", len(test.expected), len(output))
			}

			for i, str := range test.expected {
				if str != output[i] {
					t.Errorf("Incorrect FizzBuzz item in the generated list. (index: %d, expected: %s, got: %s)", i, str, output[i])
				}
			}

		})
	}
}
