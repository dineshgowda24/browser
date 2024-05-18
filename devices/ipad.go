package devices

type Ipad struct {
	p Parser
}

var (
	iPadName       = "iPad"
	iPadMatchRegex = []string{`iPad`}
)

func NewIpad(p Parser) *Ipad {
	return &Ipad{
		p: p,
	}
}

func (i *Ipad) Name() string {
	return iPadName
}

func (i *Ipad) Match() bool {
	return i.p.Match(iPadMatchRegex)
}
