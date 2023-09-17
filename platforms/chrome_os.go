package platforms

type ChromeOS struct {
	base
}

var (
	chromeOSName          = "Chrome OS"
	chromeOSVersionRegexp = []string{`CrOS [^\s]+ ([\d.]+)`}
	chromeOSMatchRegexp   = []string{`CrOS`}
)

func NewChromeOS(userAgent string) *ChromeOS {
	return &ChromeOS{
		base: base{
			userAgent: userAgent,
		},
	}
}

func (c *ChromeOS) Name() string {
	return chromeOSName
}

func (c *ChromeOS) Version() string {
	return c.version(chromeOSVersionRegexp, 1, "")
}

func (c *ChromeOS) Match() bool {
	return c.match(chromeOSMatchRegexp)
}
