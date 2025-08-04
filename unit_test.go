package promptfuse

import (
	"testing"
)

const testFile = "test_prompt.txt"

func TestAnalyze(t *testing.T) {
	cfg := Config{
		FilePath:     testFile,
		ModelName:    "gpt-4",
		PhraseLength: 2,
	}

	savingsMap, total, err := cfg.Analyze()
	if err != nil {
		t.Fatalf("Analyze() error: %v", err)
	}

	if len(savingsMap) == 0 {
		t.Errorf("Expected some savings, got none")
	}

	if total <= 0 {
		t.Errorf("Expected positive total savings, got %d", total)
	}
}

func TestBuildPhrase(t *testing.T) {
	tokens := []TokenInfo{
		{Text: "I"},
		{Text: " am"},
		{Text: " Microsoft"},
	}

	phrase, ok := BuildPhrase(tokens, 0, 3)
	if !ok {
		t.Fatalf("BuildPhrase() failed")
	}
	expected := "I am Microsoft"
	if phrase != expected {
		t.Errorf("Expected '%s', got '%s'", expected, phrase)
	}
}

func TestCountPhraseRepetition(t *testing.T) {
	tokens := []TokenInfo{
		{Text: "Go"},
		{Text: " is"},
		{Text: " fast"},
		{Text: "Go"},
		{Text: " is"},
		{Text: " fast"},
	}
	counts := CountPhraseRepetition(tokens, 3)
	if counts["Go is fast"] != 2 {
		t.Errorf("Expected 'Go is fast' to appear 2 times, got %d", counts["Go is fast"])
	}
}

func TestComputeSavingsByPhrase(t *testing.T) {
	counts := map[string]int{
		"foo bar": 6,
		"baz":     3,
	}
	savings := ComputeSavingsByPhrase(counts, 4)

	if savings["foo bar"] != 2 {
		t.Errorf("Expected 'foo bar' to save 2, got %d", savings["foo bar"])
	}

	if _, ok := savings["baz"]; ok {
		t.Errorf("'baz' should not be in the savings map")
	}
}

func TestComputeTotalSavings(t *testing.T) {
	counts := map[string]int{
		"repeat one": 7, // 3 saved
		"repeat two": 5, // 1 saved
		"not enough": 3, // no saving
	}
	total := ComputeTotalSavings(counts, 4)
	expected := 4
	if total != expected {
		t.Errorf("Expected total savings %d, got %d", expected, total)
	}
}
