package matchers

type Maxthon struct {
	p Parser
}

var (
	maxthonName          = "Maxthon"
	maxthonVersionRegexp = []string{`(?i)Maxthon/([\d.]+)`}
	maxthonMatchRegex    = []string{`(i?)Maxthon`}
)

func NewMaxthon(p Parser) *Maxthon {
	return &Maxthon{
		p: p,
	}
}

func (m *Maxthon) Name() string {
	return maxthonName
}

func (m *Maxthon) Version() string {
	return m.p.Version(maxthonVersionRegexp, 1)
}

func (m *Maxthon) Match() bool {
	return m.p.Match(maxthonMatchRegex)
}
