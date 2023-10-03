package devices

import (
	"regexp"
)

type Android struct {
	base
}

var (
	androidNameRegex  = `(?i)\(Linux.*?; Android.*?; ([-_a-z0-9 ]+)(?:;)? Build[^)]+\)`
	androidMatchRegex = []string{`(?i)Android`}
)

func NewAndroid(userAgent string) *Android {
	return &Android{
		base: base{
			userAgent: userAgent,
		},
	}
}

func (a Android) Name() string {
	re := regexp.MustCompile(androidNameRegex)
	matches := re.FindStringSubmatch(a.userAgent)
	if len(matches) > 1 {
		return matches[1]
	}

	return "Unknown"
}

func (a Android) Match() bool {
	return a.match(androidMatchRegex)
}
