package matchers

type Chrome struct {
	p Parser
}

var (
	chromeName          = "Chrome"
	chromeVersionRegexp = []string{`Chrome/([\d.]+)`, `CriOS/([\d.]+)`, `Safari/([\d.]+)`, `AppleWebKit/([\d.]+)`}
	chromeMatchRegex    = []string{`Chrome|CriOS`}
)

func NewChrome(p Parser) *Chrome {
	return &Chrome{
		p: p,
	}
}

func (c *Chrome) Name() string {
	return chromeName
}

func (c *Chrome) Version() string {
	return c.p.Version(chromeVersionRegexp, 1)
}

func (c *Chrome) Match() bool {
	return c.p.Match(chromeMatchRegex)
}
