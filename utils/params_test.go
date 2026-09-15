package utils

import "testing"

func TestParseID(t *testing.T) {
	valid := map[string]uint{"1": 1, "42": 42}
	for input, want := range valid {
		got, err := ParseID(input)
		if err != nil || got != want {
			t.Fatalf("ParseID(%q) = %d, %v; want %d", input, got, err, want)
		}
	}

	invalid := []string{"", "0", "-1", "abc", "1 OR 1=1", "(SELECT 1)", "1.5"}
	for _, input := range invalid {
		if _, err := ParseID(input); err == nil {
			t.Fatalf("ParseID(%q) should fail", input)
		}
	}
}
