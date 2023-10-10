package devices

import (
	"regexp"
)

type Samsung struct {
	base
	matches []string
}

// NewSamsung creates a new Samsung platform.
func NewSamsung(userAgent string) *Samsung {
	return &Samsung{
		base: base{
			userAgent: userAgent,
		},
		matches: make([]string, 0),
	}
}

var (
	samsungName       = "Samsung"
	samsungMatchRegex = `(?i)(?:Linux.*?; Android.*?; (?:SAMSUNG )?(SM-[A-Z0-9]+).*?)`
)

func (s *Samsung) Name() string {
	s.registerMatches()
	if len(s.matches) > 1 {
		return samsungName + " " + s.matches[1]
	}

	return samsungName
}

func (s *Samsung) Match() bool {
	s.registerMatches()
	return len(s.matches) > 0
}

func (s *Samsung) registerMatches() {
	if len(s.matches) > 0 {
		return
	}

	re := regexp.MustCompile(samsungMatchRegex)
	matches := re.FindStringSubmatch(s.userAgent)
	if len(matches) > 0 {
		s.matches = matches
	}
}
