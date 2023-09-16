package matchers

type Nokia struct {
	base
}

var (
	nokiaName          = "Nokia S40 Ovi Browser"
	nokiaVersionRegexp = []string{`S40OviBrowser/([\d.]+)`, `NokiaBrowser/([\d.]+)`}
	nokiaMatchRegexp   = []string{`S40OviBrowser`, `NokiaBrowser`}
)

func NewNokia(userAgent string) *Nokia {
	return &Nokia{
		base: newBase(userAgent),
	}
}

func (n *Nokia) Name() string {
	return nokiaName
}

func (n *Nokia) Version() string {
	return n.version(nokiaVersionRegexp, 1)
}

func (n *Nokia) Match() bool {
	return n.match(nokiaMatchRegexp)
}
