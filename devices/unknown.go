package devices

// Unknown represents an unknown device.
// These are devices that we don't know about yet. These are matched last.
type Unknown struct {
	base
}

var (
	unknownName = "Unknown"
)

func NewUnknown(userAgent string) *Unknown {
	return &Unknown{
		base: base{
			userAgent: userAgent,
		},
	}
}

func (u Unknown) Name() string {
	return unknownName
}

func (u Unknown) Match() bool {
	return true
}
