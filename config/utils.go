package config

import (
	"strings"

	"github.com/betterleaks/betterleaks/regexp"
)

func anyRegexMatch(f string, res []*regexp.Regexp) bool {
	for _, re := range res {
		if regexMatched(f, re) {
			return true
		}
	}
	return false
}

func regexMatched(f string, re *regexp.Regexp) bool {
	if re == nil {
		return false
	}
	// Use MatchString rather than FindString != "": a valid zero-width match
	// (e.g. an allowlist pattern that can match the empty string) returns "" from
	// FindString and would otherwise be wrongly treated as no match.
	return re.MatchString(f)
}

// joinRegexOr combines multiple |patterns| into a single *regexp.Regexp.
func joinRegexOr(patterns []*regexp.Regexp) *regexp.Regexp {
	var sb strings.Builder
	sb.WriteString("(?:")
	for i, pat := range patterns {
		sb.WriteString(pat.String())
		if i != len(patterns)-1 {
			sb.WriteString("|")
		}
	}
	sb.WriteString(")")
	return regexp.MustCompile(sb.String())
}
