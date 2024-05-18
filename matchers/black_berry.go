package matchers

type BlackBerry struct {
	p Parser
}

var (
	blackBerryName          = "BlackBerry"
	blackBerryVersionRegexp = []string{`BlackBerry[\da-z]+/([\d.]+)`, `Version/([\d.]+)`}
	blackBerryMatchRegex    = []string{`BlackBerry|(?i)BB10`}
)

func NewBlackBerry(p Parser) *BlackBerry {
	return &BlackBerry{
		p: p,
	}
}

func (b *BlackBerry) Name() string {
	return blackBerryName
}

func (b *BlackBerry) Version() string {
	return b.p.Version(blackBerryVersionRegexp, 1)
}

func (b *BlackBerry) Match() bool {
	return b.p.Match(blackBerryMatchRegex)
}
