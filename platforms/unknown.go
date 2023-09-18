package platforms

// Unknown is the platform information when the user agent is not recognized.
// It is used as a fallback when the user agent is not recognized.
// Its last in the list of platform matchers.
type Unknown struct {
	base
}

var (
	unknownName = "Unknown"
)

func NewUnknown(userAgent string) *Unknown {
	return &Unknown{
		base{
			userAgent: userAgent,
		},
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
