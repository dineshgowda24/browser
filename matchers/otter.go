package matchers

type Otter struct {
	p Parser
}

var (
	otterName          = "Otter"
	otterVersionRegexp = []string{`Otter/([\d.]+)`}
	otterMatchRegexp   = []string{`Otter`}
)

func NewOtter(p Parser) *Otter {
	return &Otter{
		p: p,
	}
}

func (o *Otter) Name() string {
	return otterName
}

func (o *Otter) Version() string {
	return o.p.Version(otterVersionRegexp, 1)
}

func (o *Otter) Match() bool {
	return o.p.Match(otterMatchRegexp)
}
