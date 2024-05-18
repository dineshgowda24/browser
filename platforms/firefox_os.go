package platforms

type FirefoxOS struct {
	p Parser
}

var (
	firefoxOSName              = "Firefox OS"
	firefoxOSMatchRegexp       = []string{`Firefox`}
	firefoxDeviceExcludeRegexp = []string{`Android|Linux|BlackBerry|Windows|Mac`}
)

func NewFirefoxOS(p Parser) *FirefoxOS {
	return &FirefoxOS{
		p: p,
	}
}

func (f *FirefoxOS) Name() string {
	return firefoxOSName
}

func (f *FirefoxOS) Version() string {
	return "0"
}

func (f *FirefoxOS) Match() bool {
	return !f.isExcludeDevice() && f.p.Match(firefoxOSMatchRegexp)
}

func (f *FirefoxOS) isExcludeDevice() bool {
	return f.p.Match(firefoxDeviceExcludeRegexp)
}
