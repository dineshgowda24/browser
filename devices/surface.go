package devices

import "strings"

type Surface struct {
	p Parser
}

var (
	surfaceName           = "Surface"
	surfaceWindowsRTRegex = []string{`Windows NT\s*(6\.[2-3]); ARM;`} // Windows RT 8.0 and 8.1
)

func NewSurface(p Parser) *Surface {
	return &Surface{
		p: p,
	}
}

func (s *Surface) Name() string {
	return surfaceName
}

func (s *Surface) Match() bool {
	// Matches Touch and Windows RT
	return strings.Contains(s.p.String(), "Touch") && s.p.Match(surfaceWindowsRTRegex)
}
