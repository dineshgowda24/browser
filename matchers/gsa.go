package matchers

type GoogleSearchApp struct {
	base
}

var (
	googleSearchAppName = "Google Search App"
	// TODO: review this regex, prev `GSA/([\d.]+)`
	googleSearchAppVersionRegexp = []string{`GSA/([\d]+(?:\.[\d]+)*)`}
	googleSearchAppMatchRegex    = []string{`(?i)GSA`}
)

func NewGoogleSearchApp(userAgent string) *GoogleSearchApp {
	return &GoogleSearchApp{
		base: newBase(userAgent),
	}
}

func (g *GoogleSearchApp) Name() string {
	return googleSearchAppName
}

func (g *GoogleSearchApp) Version() string {
	return g.version(googleSearchAppVersionRegexp, 1)
}

func (g *GoogleSearchApp) Match() bool {
	return g.match(googleSearchAppMatchRegex)
}
