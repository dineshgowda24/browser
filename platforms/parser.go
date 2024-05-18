package platforms

import "regexp"

type Parser interface {
	Version([]string, int, string) string
	Match([]string) bool
	String() string
}

type UAParser struct {
	userAgent string
}

// NewUAParser returns a new UAParser.
func NewUAParser(userAgent string) *UAParser {
	return &UAParser{
		userAgent: userAgent,
	}
}

func (b UAParser) String() string {
	return b.userAgent
}

// Version returns the version of the platform.
// The pattern is a list of regular expressions.
// The version is extracted from the user agent string using the given patterns.
// The order parameter specifies which match to return.
func (b UAParser) Version(patterns []string, order int, defaultVersion string) string {
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(b.userAgent)
		if len(matches) > order {
			return matches[order]
		}
	}
	return defaultVersion
}

// Match returns true if the user agent matches the pattern.
// The pattern is a list of regular expressions.
func (b UAParser) Match(patterns []string) bool {
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		if re.MatchString(b.userAgent) {
			return true
		}
	}
	return false
}
