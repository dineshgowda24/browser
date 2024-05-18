package matchers

var (
	vivaldiName          = "Vivaldi"
	vivaldiVersionRegexp = []string{`Vivaldi/([\d.]+)`}
	vivaldiMatchRegex    = []string{`(?i)Vivaldi`}
)

type Vivaldi struct {
	p Parser
}

func NewVivaldi(p Parser) *Vivaldi {
	return &Vivaldi{
		p: p,
	}
}

func (v *Vivaldi) Name() string {
	return vivaldiName
}

func (v *Vivaldi) Version() string {
	return v.p.Version(vivaldiVersionRegexp, 1)
}

func (v *Vivaldi) Match() bool {
	return v.p.Match(vivaldiMatchRegex)
}
