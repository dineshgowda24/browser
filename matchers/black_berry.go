package matchers

type BlackBerry struct {
	base
}

var (
	blackBerryName          = "BlackBerry"
	blackBerryVersionRegexp = []string{`BlackBerry[\da-z]+/([\d.]+)`, `Version/([\d.]+)`}
	blackBerryMatchRegex    = []string{`BlackBerry|(?i)BB10`}
)

func NewBlackBerry(userAgent string) *BlackBerry {
	return &BlackBerry{
		base: newBase(userAgent),
	}
}

func (b *BlackBerry) Name() string {
	return blackBerryName
}

func (b *BlackBerry) Version() string {
	return b.version(blackBerryVersionRegexp, 1)
}

func (b *BlackBerry) Match() bool {
	return b.match(blackBerryMatchRegex)
}
