package devices

type Xbox360 struct {
	p Parser
}

var (
	xbox360Name       = "Xbox 360"
	xbox360MatchRegex = []string{`(?i)Xbox`}
)

func NewXbox360(p Parser) *Xbox360 {
	return &Xbox360{
		p: p,
	}
}

func (x *Xbox360) Name() string {
	return xbox360Name
}

func (x *Xbox360) Match() bool {
	return x.p.Match(xbox360MatchRegex)
}
