package matchers

type Chrome struct {
	base
}

var (
	chromeName          = "Chrome"
	chromeVersionRegexp = []string{`Chrome/([\d.]+)`, `CriOS/([\d.]+)`, `Safari/([\d.]+)`, `AppleWebKit/([\d.]+)`}
	chromeMatchRegex    = []string{`Chrome|CriOS`}
)

func NewChrome(userAgent string) *Chrome {
	return &Chrome{
		base: newBase(userAgent),
	}
}

func (c *Chrome) Name() string {
	return chromeName
}

func (c *Chrome) Version() string {
	return c.version(chromeVersionRegexp, 1)
}

func (c *Chrome) Match() bool {
	return c.match(chromeMatchRegex)
}
