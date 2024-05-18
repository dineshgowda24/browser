package matchers

type VivoBrowser struct {
	p Parser
}

var (
	vivoBrowserName       = "Vivo Browser"
	vivoBrowserVersionReg = []string{`(?i)VivoBrowser/([\d.]+)`}
	vivoBrowserMatchRegex = []string{`(?i)VivoBrowser`}
)

func NewVivoBrowser(p Parser) *VivoBrowser {
	return &VivoBrowser{
		p: p,
	}
}

func (v *VivoBrowser) Name() string {
	return vivoBrowserName
}

func (v *VivoBrowser) Version() string {
	return v.p.Version(vivoBrowserVersionReg, 1)
}

func (v *VivoBrowser) Match() bool {
	return v.p.Match(vivoBrowserMatchRegex)
}
