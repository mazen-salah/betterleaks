package config

import (
	"testing"

	"github.com/betterleaks/betterleaks/regexp"
)

func TestRegexMatched(t *testing.T) {
	// `x*` matches the empty string at position 0, so FindString returns "" even
	// though the pattern matches. The old FindString != "" check missed this.
	if !regexMatched("yyy", regexp.MustCompile(`x*`)) {
		t.Error(`regexMatched should report a zero-width match for x* against "yyy"`)
	}
	if !regexMatched("hello world", regexp.MustCompile(`world`)) {
		t.Error("regexMatched should match a normal substring")
	}
	if regexMatched("xyz", regexp.MustCompile(`abc`)) {
		t.Error("regexMatched should be false when the pattern does not match")
	}
	if regexMatched("anything", nil) {
		t.Error("a nil regex should never match")
	}
}

func TestAnyRegexMatch(t *testing.T) {
	res := []*regexp.Regexp{regexp.MustCompile(`abc`), regexp.MustCompile(`x*`)}
	if !anyRegexMatch("zzz", res) {
		t.Error("anyRegexMatch should be true when any pattern (incl. zero-width) matches")
	}
	if anyRegexMatch("zzz", []*regexp.Regexp{regexp.MustCompile(`abc`)}) {
		t.Error("anyRegexMatch should be false when no pattern matches")
	}
}
