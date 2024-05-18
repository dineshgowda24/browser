package devices

import "regexp"

type Parser interface {
	Match([]string) bool
	String() string
}

// compile time check if UAParser implements Parser interface
var _ Parser = (*UAParser)(nil)

type UAParser struct {
	userAgent string
}

func NewUAParser(userAgent string) *UAParser {
	return &UAParser{
		userAgent: userAgent,
	}
}

func (u UAParser) String() string {
	return u.userAgent
}

// match returns true if the user agent matches the pattern.
// The pattern is a list of regular expressions.
func (u UAParser) Match(pattern []string) bool {
	for _, p := range pattern {
		re := regexp.MustCompile(p)
		if re.MatchString(u.userAgent) {
			return true
		}
	}

	return false
}
