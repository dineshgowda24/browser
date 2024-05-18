package matchers

type Safari struct {
	p Parser
}

var (
	safariName          = "Safari"
	safariVersionRegexp = []string{`Version/([\d.]+)`, `Safari/([\d.]+)`, `AppleWebKit/([\d.]+)`}
	safariMatchRegex    = []string{`Safari`}
)

func NewSafari(p Parser) *Safari {
	return &Safari{
		p: p,
	}
}

func (s *Safari) Name() string {
	return safariName
}

func (s *Safari) Version() string {
	return s.p.Version(safariVersionRegexp, 1)
}

func (s *Safari) Match() bool {
	return s.p.Match(safariMatchRegex)
}
