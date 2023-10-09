package matchers

import "regexp"

type Unknown struct {
	base
}

var (
	unknownName          = "Unknown Browser"
	unknownVersionRegexp = []string{`QuickTime/([\d.]+)`, `CoreMedia v([\d.]+)`, `AppleCoreMedia/([\d.]+)`}
)

func NewUnknown(userAgent string) *Unknown {
	return &Unknown{
		base: newBase(userAgent),
	}
}

func (u *Unknown) Name() string {
	inferedUnknowns := map[string]string{"QuickTime": "QuickTime", "CoreMedia": "Apple CoreMedia"}
	for k, v := range inferedUnknowns {
		if regexp.MustCompile(k).MatchString(u.userAgent) {
			return v
		}
	}
	return unknownName
}

func (u *Unknown) Version() string {
	return u.version(unknownVersionRegexp, 1)
}

func (u *Unknown) Match() bool {
	return true
}
