package matchers

type Instagram struct {
	base
}

var (
	instagramName          = "Instagram"
	instagramVersionRegexp = []string{`Instagram[ /]([\d.]+)`}
	instagramMatchRegex    = []string{`(?i)Instagram`}
)

func NewInstagram(userAgent string) *Instagram {
	return &Instagram{
		base: newBase(userAgent),
	}
}

func (i *Instagram) Name() string {
	return instagramName
}

func (i *Instagram) Version() string {
	return i.version(instagramVersionRegexp, 1)
}

func (i *Instagram) Match() bool {
	return i.match(instagramMatchRegex)
}
