package tptnmatch

import "testing"

func TestCanMatchUnicode(t *testing.T) {
	var matchedPatterns []string

	MatchTextAgainstTrie("This program 💻 is as cool 😎 as 🔥fighter and can match unicode patterns like a boss 🧨😎🚀",
		BuildTrie([]string{"as", "😎", "🦑", "🔥fighter"}), func(p string) {
			matchedPatterns = append(matchedPatterns, p)
		})

	expectedMatches := []string{"as", "😎", "as", "🔥fighter", "😎"}

	if len(expectedMatches) != len(matchedPatterns) {
		t.Fatalf("Expected %d matches got %d\n", len(expectedMatches), len(matchedPatterns))
	}

	for idx := 0; idx < len(matchedPatterns); idx++ {
		if matchedPatterns[idx] != expectedMatches[idx] {
			t.Errorf("%d match expected to be '%s', got '%s'", idx, expectedMatches[idx], matchedPatterns[idx])
		}
	}
}

// This test just fixes and makes clear current limitation.
func TestCanMatchOverlappingPrefixes(t *testing.T) {
	var matchedPatterns []string

	MatchTextAgainstTrie("🔥 neutralize by 🔥fighter 🔥f",
		BuildTrie([]string{"🔥", "🔥fighter"}), func(p string) {
			matchedPatterns = append(matchedPatterns, p)
		})

	expectedMatches := []string{"🔥", "🔥fighter"} // 🔥f is partial prefix and should not match.

	if len(expectedMatches) != len(matchedPatterns) {
		t.Fatalf("Expected %d matches got %d\n", len(expectedMatches), len(matchedPatterns))
	}

	for idx := 0; idx < len(matchedPatterns); idx++ {
		if matchedPatterns[idx] != expectedMatches[idx] {
			t.Errorf("%d match expected to be '%s', got '%s'", idx, expectedMatches[idx], matchedPatterns[idx])
		}
	}
}
