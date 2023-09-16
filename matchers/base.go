package matchers

import (
	"regexp"
)

type base struct {
	userAgent string
}

func newBase(userAgent string) base {
	return base{
		userAgent: userAgent,
	}
}

// version returns the first match of the given patterns.
// The pattern is a list of regular expressions.
// The order is the index of the match group in the regular expression.
// If the order is greater than the number of matches, it returns "0.0".
func (b base) version(patterns []string, order int) string {
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(b.userAgent)

		if len(matches) > order {
			return matches[order]
		}
	}
	return "0.0"
}

// match returns true if the user agent matches the pattern.
// The pattern is a list of regular expressions.
func (b base) match(patterns []string) bool {
	for _, pattern := range patterns {
		if regexp.MustCompile(pattern).MatchString(b.userAgent) {
			return true
		}
	}
	return false
}
