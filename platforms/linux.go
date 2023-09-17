package platforms

type Linux struct {
	base
}

var (
	linuxName        = "Generic Linux"
	linuxMatchRegexp = []string{`Linux`}
)

func NewLinux(userAgent string) *Linux {
	return &Linux{
		base: base{
			userAgent: userAgent,
		},
	}
}

func (l *Linux) Name() string {
	return linuxName
}

func (l *Linux) Version() string {
	return "0"
}

func (l *Linux) Match() bool {
	return l.match(linuxMatchRegexp)
}
