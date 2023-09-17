package matchers

type SamsungBrowser struct {
	base
}

var (
	samsungBrowserName         = "Samsung Browser"
	samsungBrowserVersionRegex = []string{
		`(?i)SamsungBrowser/([\d.]+)`,
		`Chrome/([\d.]+)`,
		`CriOS/([\d.]+)`,
		`Safari/([\d.]+)`,
		`AppleWebKit/([\d.]+)`,
	}
	samsungBrowserMatchRegex = []string{`SamsungBrowser`}
)

func NewSamsungBrowser(userAgent string) *SamsungBrowser {
	return &SamsungBrowser{
		base{
			userAgent: userAgent,
		},
	}
}

func (s *SamsungBrowser) Name() string {
	return samsungBrowserName
}

func (s *SamsungBrowser) Version() string {
	return s.version(samsungBrowserVersionRegex, 1)
}

func (s *SamsungBrowser) Match() bool {
	return s.match(samsungBrowserMatchRegex)
}
