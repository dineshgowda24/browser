package matchers

type SogouBrowser struct {
	p Parser
}

var (
	sogouBrowserName          = "Sogou Browser"
	sogouBrowserVersionRegexp = []string{`(?i)(?:SogouMobileBrowser)/([\d.]+)`}
	sogouBrowserMatchRegexp   = []string{`(?i)SogouMobileBrowser`, `\bSE\b`}
)

func NewSogouBrowser(p Parser) *SogouBrowser {
	return &SogouBrowser{
		p: p,
	}
}

func (s *SogouBrowser) Name() string {
	return sogouBrowserName
}

func (s *SogouBrowser) Version() string {
	return s.p.Version(sogouBrowserVersionRegexp, 1)
}

func (s *SogouBrowser) Match() bool {
	return s.p.Match(sogouBrowserMatchRegexp)
}
