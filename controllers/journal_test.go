package controllers

import "testing"

func TestCleanTextEncodesForPDFCoreFonts(t *testing.T) {
	tests := map[string]string{
		"plain text":      "plain text",
		"café naïve Zoë":  "caf\xe9 na\xefve Zo\xeb",
		"€ 5":             "\x80 5",
		"“quoted” – it’s": "\"quoted\" - it's",
		"non breaking":    "non breaking",
		"日本 🙂":            "?? ?",
	}
	for input, want := range tests {
		if got := cleanText(input); got != want {
			t.Errorf("cleanText(%q) = %q, want %q", input, got, want)
		}
	}
}
