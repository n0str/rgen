package main

import (
	"strings"
	"testing"
)

// Characters that are easy to misread on a screen, see lettersSafe.
const lookalikes = "0OoQ1Il2Zz5Ss8BgquvVr"

// Letters whose upper and lower cases differ only in size.
const caseLookalikes = "CcKkOoPpSsUuVvWwXxYyZz"

func TestSafeAlphabets(t *testing.T) {
	alphabets := map[string][]rune{
		"idsafe":  lettersSafe,
		"idsafel": lettersSafeLower,
		"idsafeu": lettersSafeUpper,
	}
	for name, letters := range alphabets {
		seen := map[rune]bool{}
		for _, r := range letters {
			if seen[r] {
				t.Errorf("%s: duplicate %q", name, r)
			}
			seen[r] = true
			if strings.ContainsRune(lookalikes, r) {
				t.Errorf("%s: lookalike %q", name, r)
			}
		}
	}

	for _, r := range lettersSafe {
		if strings.ContainsRune(caseLookalikes, r) {
			t.Errorf("idsafe: %q differs from its other case only in size", r)
		}
	}
}

func TestRandSeqUsesAlphabet(t *testing.T) {
	for _, letters := range [][]rune{lettersSafe, lettersSafeLower, lettersSafeUpper} {
		for _, r := range randSeq(1000, letters) {
			if !strings.ContainsRune(string(letters), r) {
				t.Fatalf("randSeq produced %q outside of %q", r, string(letters))
			}
		}
	}
}
