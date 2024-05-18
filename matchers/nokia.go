package matchers

type Nokia struct {
	p Parser
}

var (
	nokiaName          = "Nokia S40 Ovi Browser"
	nokiaVersionRegexp = []string{`S40OviBrowser/([\d.]+)`, `NokiaBrowser/([\d.]+)`}
	nokiaMatchRegexp   = []string{`S40OviBrowser`, `NokiaBrowser`}
)

func NewNokia(p Parser) *Nokia {
	return &Nokia{
		p: p,
	}
}

func (n *Nokia) Name() string {
	return nokiaName
}

func (n *Nokia) Version() string {
	return n.p.Version(nokiaVersionRegexp, 1)
}

func (n *Nokia) Match() bool {
	return n.p.Match(nokiaMatchRegexp)
}
