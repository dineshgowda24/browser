package matchers

type VivoBrowser struct {
	base
}

var (
	vivoBrowserName       = "Vivo Browser"
	vivoBrowserVersionReg = []string{`(?i)VivoBrowser/([\d.]+)`}
	vivoBrowserMatchRegex = []string{`(?i)VivoBrowser`}
)

func NewVivoBrowser(userAgent string) *VivoBrowser {
	return &VivoBrowser{
		base{
			userAgent: userAgent,
		},
	}
}

func (v *VivoBrowser) Name() string {
	return vivoBrowserName
}

func (v *VivoBrowser) Version() string {
	return v.version(vivoBrowserVersionReg, 1)
}

func (v *VivoBrowser) Match() bool {
	return v.match(vivoBrowserMatchRegex)
}
