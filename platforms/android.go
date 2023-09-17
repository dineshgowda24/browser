package platforms

type Android struct {
	base
}

var (
	androidName          = "Android"
	androidVersionRegexp = []string{`(?i)Android ([\d.]+)`}
	androidMatchRegexp   = []string{`(?i)Android`}
)

func NewAndroid(userAgent string) *Android {
	return &Android{
		base: base{
			userAgent: userAgent,
		},
	}
}

func (a *Android) Name() string {
	return androidName
}

func (a *Android) Version() string {
	return a.version(androidVersionRegexp, 1, "")
}

func (a *Android) Match() bool {
	return a.match(androidMatchRegexp) && !a.match([]string{`KAIOS`})
}
