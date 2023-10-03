package matchers

type Safari struct {
	base
}

var (
	safariName          = "Safari"
	safariVersionRegexp = []string{`Version/([\d.]+)`, `Safari/([\d.]+)`, `AppleWebKit/([\d.]+)`}
	safariMatchRegex    = []string{`Safari`}
)

func NewSafari(userAgent string) *Safari {
	return &Safari{
		base: newBase(userAgent),
	}
}

func (s *Safari) Name() string {
	return safariName
}

func (s *Safari) Version() string {
	return s.version(safariVersionRegexp, 1)
}

func (s *Safari) Match() bool {
	return s.match(safariMatchRegex)
}
