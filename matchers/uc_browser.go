package matchers

type UCBrowser struct {
	base
}

var (
	ucBrowserName          = "UCBrowser"
	ucBrowserVersionRegexp = []string{`UCBrowser/([\d.]+)`, `UCWEB(?:/)?([\d.]+)`}
	ucBrowserMatchRegexp   = []string{`UC(Browser|WEB)`}
)

func NewUCBrowser(userAgent string) *UCBrowser {
	return &UCBrowser{
		base: newBase(userAgent),
	}
}

func (u *UCBrowser) Name() string {
	return ucBrowserName
}

func (u *UCBrowser) Version() string {
	return u.version(ucBrowserVersionRegexp, 1)
}

func (u *UCBrowser) Match() bool {
	return u.match(ucBrowserMatchRegexp)
}
