package matchers

type YaaniBrowser struct {
	base
}

var (
	yaaniBrowserName          = "Yaani Browser"
	yaaniBrowserVersionRegexp = []string{`YaaniBrowser/([\d.]+)`, `Turkcell-YaaniBrowser/([\d.]+)`, `Chrome/([\d.]+)`}
	yaaniBrowserMatchRegex    = []string{`YaaniBrowser`}
)

func NewYaaniBrowser(userAgent string) *YaaniBrowser {
	return &YaaniBrowser{
		base: newBase(userAgent),
	}
}

func (y *YaaniBrowser) Name() string {
	return yaaniBrowserName
}

func (y *YaaniBrowser) Version() string {
	return y.version(yaaniBrowserVersionRegexp, 1)
}

func (y *YaaniBrowser) Match() bool {
	return y.match(yaaniBrowserMatchRegex)
}
