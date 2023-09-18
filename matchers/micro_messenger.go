package matchers

type MicroMessenger struct {
	base
}

var (
	microMessengerName          = "MicroMessenger"
	microMessengerVersionRegexp = []string{`(?i)(?:MicroMessenger)/([\d.]+)`}
	microMessengerMatchRegexp   = []string{`MicroMessenger`}
)

func NewMicroMessenger(userAgent string) *MicroMessenger {
	return &MicroMessenger{
		base: newBase(userAgent),
	}
}

func (m *MicroMessenger) Name() string {
	return microMessengerName
}

func (m *MicroMessenger) Version() string {
	return m.version(microMessengerVersionRegexp, 1)
}

func (m *MicroMessenger) Match() bool {
	return m.match(microMessengerMatchRegexp)
}
