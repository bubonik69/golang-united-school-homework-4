package string_sum

import (
	"testing"
)

func TestStringSum(t *testing.T) {
	tests := []struct {
		value  string
		result string
	}{
		{"-3+5", "2"},
		{" - 3 -5", "-8"},
		{" 3 + 5 ", "8"},
		{"    ", ""},
	}
	for i := 0; i < len(tests); i++ {
		if s, err := StringSum(tests[i].value); s != tests[i].result {
			t.Errorf("Error in %d test value. Result is `%s`. Need `%s`", i, s, tests[i].result)
		} else {
			t.Logf("Test %d passed good, error - %e", i, err)
		}

	}
}
