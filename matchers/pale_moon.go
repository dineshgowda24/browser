package matchers

type PaleMoon struct {
	p Parser
}

var (
	paleMoonName          = "Pale Moon"
	paleMoonVersionRegexp = []string{`PaleMoon/([\d.]+)`}
	paleMoonMatchRegex    = []string{`PaleMoon`}
)

func NewPaleMoon(p Parser) *PaleMoon {
	return &PaleMoon{
		p: p,
	}
}

func (pa *PaleMoon) Name() string {
	return paleMoonName
}

func (pa *PaleMoon) Version() string {
	return pa.p.Version(paleMoonVersionRegexp, 1)
}

func (pa *PaleMoon) Match() bool {
	return pa.p.Match(paleMoonMatchRegex)
}
