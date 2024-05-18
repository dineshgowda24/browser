package matchers

type YaaniBrowser struct {
	p Parser
}

var (
	yaaniBrowserName          = "Yaani Browser"
	yaaniBrowserVersionRegexp = []string{`YaaniBrowser/([\d.]+)`, `Turkcell-YaaniBrowser/([\d.]+)`, `Chrome/([\d.]+)`}
	yaaniBrowserMatchRegex    = []string{`YaaniBrowser`}
)

func NewYaaniBrowser(p Parser) *YaaniBrowser {
	return &YaaniBrowser{
		p: p,
	}
}

func (y *YaaniBrowser) Name() string {
	return yaaniBrowserName
}

func (y *YaaniBrowser) Version() string {
	return y.p.Version(yaaniBrowserVersionRegexp, 1)
}

func (y *YaaniBrowser) Match() bool {
	return y.p.Match(yaaniBrowserMatchRegex)
}
