package matchers

type Sputnik struct {
	p Parser
}

var (
	sputnikName          = "Sputnik"
	sputnikVersionRegexp = []string{`SputnikBrowser/([\d.]+)`}
	sputnikMatchRegexp   = []string{`SputnikBrowser`}
)

func NewSputnik(p Parser) *Sputnik {
	return &Sputnik{
		p: p,
	}
}

func (s *Sputnik) Name() string {
	return sputnikName
}

func (s *Sputnik) Version() string {
	return s.p.Version(sputnikVersionRegexp, 1)
}

func (s *Sputnik) Match() bool {
	return s.p.Match(sputnikMatchRegexp)
}
