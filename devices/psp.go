package devices

type PSP struct {
	base
}

var (
	pspName       = "PlayStation Portable"
	pspMatchRegex = []string{`(?i)PlayStation Portable`}
)

func NewPSP(userAgent string) *PSP {
	return &PSP{
		base{
			userAgent: userAgent,
		},
	}
}

func (p *PSP) Name() string {
	return pspName
}

func (p *PSP) Match() bool {
	return p.match(pspMatchRegex)
}
