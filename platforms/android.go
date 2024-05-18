package platforms

type Android struct {
	p Parser
}

var (
	androidName          = "Android"
	androidVersionRegexp = []string{`(?i)Android ([\d.]+)`}
	androidMatchRegexp   = []string{`(?i)Android`}
)

func NewAndroid(p Parser) *Android {
	return &Android{
		p: p,
	}
}

func (a *Android) Name() string {
	return androidName
}

func (a *Android) Version() string {
	return a.p.Version(androidVersionRegexp, 1, "")
}

func (a *Android) Match() bool {
	return a.p.Match(androidMatchRegexp) && !a.p.Match([]string{`KAIOS`})
}
