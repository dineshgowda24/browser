package devices

import (
	"regexp"
)

type Android struct {
	p Parser
}

var (
	androidNameRegex  = `(?i)\(Linux.*?; Android.*?; ([-_a-z0-9 ]+)(?:;)? Build[^)]+\)`
	androidMatchRegex = []string{`(?i)Android`}
)

func NewAndroid(p Parser) *Android {
	return &Android{
		p: p,
	}
}

func (a Android) Name() string {
	re := regexp.MustCompile(androidNameRegex)
	matches := re.FindStringSubmatch(a.p.String())
	if len(matches) > 1 {
		return matches[1]
	}

	return "Unknown"
}

func (a Android) Match() bool {
	return a.p.Match(androidMatchRegex)
}
