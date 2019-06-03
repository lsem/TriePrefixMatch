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

	MatchTextAgainstTrie("🔥 neutralized by 🔥fighter 🔥f",
		BuildTrie([]string{"🔥", "🔥fighter", "neutralize", "neut"}), func(p string) {
			matchedPatterns = append(matchedPatterns, p)
		})

	expectedMatches := []string{"🔥", "neutralize", "🔥fighter"} // 🔥f is partial prefix and should not match.

	if len(expectedMatches) != len(matchedPatterns) {
		t.Fatalf("Expected %d matches got %d\n", len(expectedMatches), len(matchedPatterns))
	}

	for idx := 0; idx < len(matchedPatterns); idx++ {
		if matchedPatterns[idx] != expectedMatches[idx] {
			t.Errorf("%d match expected to be '%s', got '%s'", idx, expectedMatches[idx], matchedPatterns[idx])
		}
	}
}

func TestCanMatchOneCharacter(t *testing.T) {
	var matchedPatterns []string

	MatchTextAgainstTrie("A",
		BuildTrie([]string{"A"}), func(p string) {
			matchedPatterns = append(matchedPatterns, p)
		})

	expectedMatches := []string{"A"}

	if len(expectedMatches) != len(matchedPatterns) {
		t.Fatalf("Expected %d matches got %d\n", len(expectedMatches), len(matchedPatterns))
	}

	for idx := 0; idx < len(matchedPatterns); idx++ {
		if matchedPatterns[idx] != expectedMatches[idx] {
			t.Errorf("%d match expected to be '%s', got '%s'", idx, expectedMatches[idx], matchedPatterns[idx])
		}
	}
}

func TestCannotMatchWrongCharcter(t *testing.T) {
	var matchedPatterns []string

	MatchTextAgainstTrie("B",
		BuildTrie([]string{"A"}), func(p string) {
			matchedPatterns = append(matchedPatterns, p)
		})

	var expectedMatches []string

	if len(expectedMatches) != len(matchedPatterns) {
		t.Fatalf("Expected %d matches got %d\n", len(expectedMatches), len(matchedPatterns))
	}

	for idx := 0; idx < len(matchedPatterns); idx++ {
		if matchedPatterns[idx] != expectedMatches[idx] {
			t.Errorf("%d match expected to be '%s', got '%s'", idx, expectedMatches[idx], matchedPatterns[idx])
		}
	}
}


func TestCanMatchLastCharacter(t *testing.T) {
	var matchedPatterns []string

	MatchTextAgainstTrie("BCDA",
		BuildTrie([]string{"A"}), func(p string) {
			matchedPatterns = append(matchedPatterns, p)
		})

	expectedMatches := []string{"A"}

	if len(expectedMatches) != len(matchedPatterns) {
		t.Fatalf("Expected %d matches got %d\n", len(expectedMatches), len(matchedPatterns))
	}

	for idx := 0; idx < len(matchedPatterns); idx++ {
		if matchedPatterns[idx] != expectedMatches[idx] {
			t.Errorf("%d match expected to be '%s', got '%s'", idx, expectedMatches[idx], matchedPatterns[idx])
		}
	}
}

func TestCannotMatchLastCharacterOfTextAgainstInCompletePattern(t *testing.T) {
	var matchedPatterns []string

	MatchTextAgainstTrie("AB",
		BuildTrie([]string{"BX"}), func(p string) {
			matchedPatterns = append(matchedPatterns, p)
		})

	expectedMatches := []string{}

	if len(expectedMatches) != len(matchedPatterns) {
		t.Fatalf("Expected %d matches got %d\n", len(expectedMatches), len(matchedPatterns))
	}

	for idx := 0; idx < len(matchedPatterns); idx++ {
		if matchedPatterns[idx] != expectedMatches[idx] {
			t.Errorf("%d match expected to be '%s', got '%s'", idx, expectedMatches[idx], matchedPatterns[idx])
		}
	}
}
