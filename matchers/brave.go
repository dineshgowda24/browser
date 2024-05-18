package matchers

var (
	braveName          = "Brave"
	braveVersionRegexp = []string{`brave/([\d.]+)`}
	braveMatchRegex    = []string{`(?i)Brave`}
)

type Brave struct {
	p Parser
}

func NewBrave(p Parser) *Brave {
	return &Brave{
		p: p,
	}
}

func (b *Brave) Name() string {
	return braveName
}

func (b *Brave) Version() string {
	return b.p.Version(braveVersionRegexp, 1)
}

func (b *Brave) Match() bool {
	return b.p.Match(braveMatchRegex)
}
