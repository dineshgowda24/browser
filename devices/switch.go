package devices

type Switch struct {
	p Parser
}

var (
	switchName       = "Nintendo Switch"
	switchMatchRegex = []string{`(?i)Nintendo Switch`}
)

func NewSwitch(p Parser) *Switch {
	return &Switch{
		p: p,
	}
}

func (s *Switch) Name() string {
	return switchName
}

func (s *Switch) Match() bool {
	return s.p.Match(switchMatchRegex)
}
