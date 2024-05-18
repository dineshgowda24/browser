package matchers

type MicroMessenger struct {
	p Parser
}

var (
	microMessengerName          = "MicroMessenger"
	microMessengerVersionRegexp = []string{`(?i)(?:MicroMessenger)/([\d.]+)`}
	microMessengerMatchRegexp   = []string{`MicroMessenger`}
)

func NewMicroMessenger(p Parser) *MicroMessenger {
	return &MicroMessenger{
		p: p,
	}
}

func (m *MicroMessenger) Name() string {
	return microMessengerName
}

func (m *MicroMessenger) Version() string {
	return m.p.Version(microMessengerVersionRegexp, 1)
}

func (m *MicroMessenger) Match() bool {
	return m.p.Match(microMessengerMatchRegexp)
}
