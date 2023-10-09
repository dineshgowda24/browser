package devices

import "strings"

type Surface struct {
	base
}

var (
	surfaceName           = "Surface"
	surfaceWindowsRTRegex = []string{`Windows NT\s*(6\.[2-3]); ARM;`} // Windows RT 8.0 and 8.1
)

func NewSurface(userAgent string) *Surface {
	return &Surface{
		base: base{
			userAgent: userAgent,
		},
	}
}

func (s *Surface) Name() string {
	return surfaceName
}

func (s *Surface) Match() bool {
	// Matches Touch and Windows RT
	return strings.Contains(s.userAgent, "Touch") && s.match(surfaceWindowsRTRegex)
}
