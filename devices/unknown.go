package devices

// Unknown represents an unknown device.
// These are devices that we don't know about yet. These are matched last.
type Unknown struct {
	p Parser
}

var unknownName = "Unknown"

func NewUnknown(p Parser) *Unknown {
	return &Unknown{
		p: p,
	}
}

func (u Unknown) Name() string {
	return unknownName
}

func (u Unknown) Match() bool {
	return true
}
