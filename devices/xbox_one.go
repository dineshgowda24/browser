package devices

type XboxOne struct {
	p Parser
}

var (
	xboxOneName       = "Xbox One"
	xboxOneMatchRegex = []string{`(?i)Xbox One`}
)

func NewXboxOne(p Parser) *XboxOne {
	return &XboxOne{
		p: p,
	}
}

func (x *XboxOne) Name() string {
	return xboxOneName
}

func (x *XboxOne) Match() bool {
	return x.p.Match(xboxOneMatchRegex)
}
