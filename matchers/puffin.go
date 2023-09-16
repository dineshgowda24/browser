package matchers

type Puffin struct {
	base
}

var (
	puffinName       = "Puffin"
	puffinVersionReg = []string{`(?i)Puffin/([\d.]+)`}
	puffinMatchRegex = []string{`(?i)Puffin`}
)

func NewPuffin(userAgent string) *Puffin {
	return &Puffin{
		base{
			userAgent: userAgent,
		},
	}
}

func (p *Puffin) Name() string {
	return puffinName
}

func (p *Puffin) Version() string {
	return p.version(puffinVersionReg, 1)
}

func (p *Puffin) Match() bool {
	return p.match(puffinMatchRegex)
}
