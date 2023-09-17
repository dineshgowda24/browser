package platforms

import "regexp"

type base struct {
	userAgent string
}

// version returns the version of the platform.
// The pattern is a list of regular expressions.
// The version is extracted from the user agent string using the given patterns.
// The order parameter specifies which match to return.
func (b base) version(patterns []string, order int, defaultVersion string) string {
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(b.userAgent)
		if len(matches) > order {
			return matches[order]
		}
	}
	return defaultVersion
}

// match returns true if the user agent matches the pattern.
// The pattern is a list of regular expressions.
func (b base) match(patterns []string) bool {
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		if re.MatchString(b.userAgent) {
			return true
		}
	}
	return false
}
