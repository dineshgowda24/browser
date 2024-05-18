package matchers

type Snapchat struct {
	p Parser
}

var (
	snapchatName          = "Snapchat"
	snapchatVersionRegexp = []string{`Snapchat(?: |/)?([\d.]+)`}
	snapchatMatchRegexp   = []string{`Snapchat`}
)

func NewSnapchat(p Parser) *Snapchat {
	return &Snapchat{
		p: p,
	}
}

func (s *Snapchat) Name() string {
	return snapchatName
}

func (s *Snapchat) Version() string {
	return s.p.Version(snapchatVersionRegexp, 1)
}

func (s *Snapchat) Match() bool {
	return s.p.Match(snapchatMatchRegexp)
}
