package platforms

type BlackBerry struct {
	base
}

var (
	blackBerryName          = "BlackBerry"
	blackBerryVersionRegexp = []string{`(?:Version|BlackBerry[\da-z]+)/([\d.]+)`}
	blackBerryMatchRegexp   = []string{`BB10|BlackBerry`}
)

func NewBlackBerry(userAgent string) *BlackBerry {
	return &BlackBerry{
		base: base{
			userAgent: userAgent,
		},
	}
}

func (b *BlackBerry) Name() string {
	return blackBerryName
}

func (b *BlackBerry) Version() string {
	return b.version(blackBerryVersionRegexp, 1, "")
}

func (b *BlackBerry) Match() bool {
	return b.match(blackBerryMatchRegexp)
}
