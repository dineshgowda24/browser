package platforms

type Linux struct {
	p Parser
}

var (
	linuxName        = "Generic Linux"
	linuxMatchRegexp = []string{`Linux`}
)

func NewLinux(p Parser) *Linux {
	return &Linux{
		p: p,
	}
}

func (l *Linux) Name() string {
	return linuxName
}

func (l *Linux) Version() string {
	return "0"
}

func (l *Linux) Match() bool {
	return l.p.Match(linuxMatchRegexp)
}
