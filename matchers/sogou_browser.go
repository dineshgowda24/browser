package matchers

type SogouBrowser struct {
	base
}

var (
	sogouBrowserName          = "Sogou Browser"
	sogouBrowserVersionRegexp = []string{`(?i)(?:SogouMobileBrowser)/([\d.]+)`}
	sogouBrowserMatchRegexp   = []string{`(?i)SogouMobileBrowser`, `\bSE\b`}
)

func NewSogouBrowser(userAgent string) *SogouBrowser {
	return &SogouBrowser{
		base: newBase(userAgent),
	}
}

func (s *SogouBrowser) Name() string {
	return sogouBrowserName
}

func (s *SogouBrowser) Version() string {
	return s.version(sogouBrowserVersionRegexp, 1)
}

func (s *SogouBrowser) Match() bool {
	return s.match(sogouBrowserMatchRegexp)
}
