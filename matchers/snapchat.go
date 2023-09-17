package matchers

type Snapchat struct {
	base
}

var (
	snapchatName          = "Snapchat"
	snapchatVersionRegexp = []string{`Snapchat(?: |/)?([\d.]+)`}
	snapchatMatchRegexp   = []string{`Snapchat`}
)

func NewSnapchat(userAgent string) *Snapchat {
	return &Snapchat{
		base: newBase(userAgent),
	}
}

func (s *Snapchat) Name() string {
	return snapchatName
}

func (s *Snapchat) Version() string {
	return s.version(snapchatVersionRegexp, 1)
}

func (s *Snapchat) Match() bool {
	return s.match(snapchatMatchRegexp)
}
