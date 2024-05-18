package matchers

type Instagram struct {
	p Parser
}

var (
	instagramName          = "Instagram"
	instagramVersionRegexp = []string{`Instagram[ /]([\d.]+)`}
	instagramMatchRegex    = []string{`(?i)Instagram`}
)

func NewInstagram(p Parser) *Instagram {
	return &Instagram{
		p: p,
	}
}

func (i *Instagram) Name() string {
	return instagramName
}

func (i *Instagram) Version() string {
	return i.p.Version(instagramVersionRegexp, 1)
}

func (i *Instagram) Match() bool {
	return i.p.Match(instagramMatchRegex)
}
