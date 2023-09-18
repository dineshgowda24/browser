package devices

import "regexp"

type base struct {
	userAgent string
}

// match returns true if the user agent matches the pattern.
// The pattern is a list of regular expressions.
func (b *base) match(pattern []string) bool {
	for _, p := range pattern {
		re := regexp.MustCompile(p)
		if re.MatchString(b.userAgent) {
			return true
		}
	}

	return false
}
