package devices

type Switch struct {
	base
}

var (
	switchName       = "Nintendo Switch"
	switchMatchRegex = []string{`(?i)Nintendo Switch`}
)

func NewSwitch(userAgent string) *Switch {
	return &Switch{
		base{
			userAgent: userAgent,
		},
	}
}

func (s *Switch) Name() string {
	return switchName
}

func (s *Switch) Match() bool {
	return s.match(switchMatchRegex)
}
