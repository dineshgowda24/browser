package platforms

// Unknown is the platform information when the user agent is not recognized.
// It is used as a fallback when the user agent is not recognized.
// Its last in the list of platform matchers.
type Unknown struct {
	p Parser
}

var unknownName = "Unknown"

func NewUnknown(p Parser) *Unknown {
	return &Unknown{
		p: p,
	}
}

func (u *Unknown) Name() string {
	return unknownName
}

func (u *Unknown) Version() string {
	return "0"
}

func (u *Unknown) Match() bool {
	return true
}
