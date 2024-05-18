package matchers

import "regexp"

type Unknown struct {
	p Parser
}

var (
	unknownName          = "Unknown Browser"
	unknownVersionRegexp = []string{`QuickTime/([\d.]+)`, `CoreMedia v([\d.]+)`, `AppleCoreMedia/([\d.]+)`}
)

func NewUnknown(p Parser) *Unknown {
	return &Unknown{
		p: p,
	}
}

func (u *Unknown) Name() string {
	inferedUnknowns := map[string]string{"QuickTime": "QuickTime", "CoreMedia": "Apple CoreMedia"}
	for k, v := range inferedUnknowns {
		if regexp.MustCompile(k).MatchString(u.p.String()) {
			return v
		}
	}
	return unknownName
}

func (u *Unknown) Version() string {
	return u.p.Version(unknownVersionRegexp, 1)
}

func (u *Unknown) Match() bool {
	return true
}
