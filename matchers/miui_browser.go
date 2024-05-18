package matchers

type MiuiBrowser struct {
	p Parser
}

var (
	miuiBrowserName          = "Miui Browser"
	miuiBrowserVersionRegexp = []string{`MiuiBrowser/([\d.]+)`}
	miuiBrowserMatchRegexp   = []string{`MiuiBrowser`}
)

func NewMiuiBrowser(p Parser) *MiuiBrowser {
	return &MiuiBrowser{
		p: p,
	}
}

func (m *MiuiBrowser) Name() string {
	return miuiBrowserName
}

func (m *MiuiBrowser) Version() string {
	return m.p.Version(miuiBrowserVersionRegexp, 1)
}

func (m *MiuiBrowser) Match() bool {
	return m.p.Match(miuiBrowserMatchRegexp)
}
