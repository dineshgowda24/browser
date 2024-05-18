package matchers

type UCBrowser struct {
	p Parser
}

var (
	ucBrowserName          = "UCBrowser"
	ucBrowserVersionRegexp = []string{`UCBrowser/([\d.]+)`, `UCWEB(?:/)?([\d.]+)`}
	ucBrowserMatchRegexp   = []string{`UC(Browser|WEB)`}
)

func NewUCBrowser(p Parser) *UCBrowser {
	return &UCBrowser{
		p: p,
	}
}

func (u *UCBrowser) Name() string {
	return ucBrowserName
}

func (u *UCBrowser) Version() string {
	return u.p.Version(ucBrowserVersionRegexp, 1)
}

func (u *UCBrowser) Match() bool {
	return u.p.Match(ucBrowserMatchRegexp)
}
