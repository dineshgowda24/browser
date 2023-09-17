package platforms

type FirefoxOS struct {
	base
}

var (
	firefoxOSName              = "Firefox OS"
	firefoxOSMatchRegexp       = []string{`Firefox`}
	firefoxDeviceExcludeRegexp = []string{`Android|Linux|BlackBerry|Windows|Mac`}
)

func NewFirefoxOS(userAgent string) *FirefoxOS {
	return &FirefoxOS{
		base{
			userAgent: userAgent,
		},
	}
}

func (f *FirefoxOS) Name() string {
	return firefoxOSName
}

func (f *FirefoxOS) Version() string {
	return "0"
}

func (f *FirefoxOS) Match() bool {
	return !f.isExcludeDevice() && f.match(firefoxOSMatchRegexp)
}

func (f *FirefoxOS) isExcludeDevice() bool {
	return f.match(firefoxDeviceExcludeRegexp)
}
