package matchers

type Puffin struct {
	p Parser
}

var (
	puffinName       = "Puffin"
	puffinVersionReg = []string{`(?i)Puffin/([\d.]+)`}
	puffinMatchRegex = []string{`(?i)Puffin`}
)

func NewPuffin(p Parser) *Puffin {
	return &Puffin{
		p: p,
	}
}

func (pu *Puffin) Name() string {
	return puffinName
}

func (pu *Puffin) Version() string {
	return pu.p.Version(puffinVersionReg, 1)
}

func (pu *Puffin) Match() bool {
	return pu.p.Match(puffinMatchRegex)
}
