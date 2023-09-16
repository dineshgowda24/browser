package matchers

type Maxthon struct {
	base
}

var (
	maxthonName          = "Maxthon"
	maxthonVersionRegexp = []string{`(?i)Maxthon/([\d.]+)`}
	maxthonMatchRegex    = []string{`(i?)Maxthon`}
)

func NewMaxthon(userAgent string) *Maxthon {
	return &Maxthon{
		base: newBase(userAgent),
	}
}

func (m *Maxthon) Name() string {
	return maxthonName
}

func (m *Maxthon) Version() string {
	return m.version(maxthonVersionRegexp, 1)
}

func (m *Maxthon) Match() bool {
	return m.match(maxthonMatchRegex)
}
