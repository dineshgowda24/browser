package matchers

type PaleMoon struct {
	base
}

var (
	paleMoonName          = "Pale Moon"
	paleMoonVersionRegexp = []string{`PaleMoon/([\d.]+)`}
	paleMoonMatchRegex    = []string{`PaleMoon`}
)

func NewPaleMoon(userAgent string) *PaleMoon {
	return &PaleMoon{
		base: newBase(userAgent),
	}
}

func (p *PaleMoon) Name() string {
	return paleMoonName
}

func (p *PaleMoon) Version() string {
	return p.version(paleMoonVersionRegexp, 1)
}

func (p *PaleMoon) Match() bool {
	return p.match(paleMoonMatchRegex)
}
