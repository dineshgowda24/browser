package devices

type PSP struct {
	p Parser
}

var (
	pspName       = "PlayStation Portable"
	pspMatchRegex = []string{`(?i)PlayStation Portable`}
)

func NewPSP(p Parser) *PSP {
	return &PSP{
		p: p,
	}
}

func (p *PSP) Name() string {
	return pspName
}

func (p *PSP) Match() bool {
	return p.p.Match(pspMatchRegex)
}
