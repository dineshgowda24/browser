package devices

type PlayStation5 struct {
	base
}

var (
	playStation5Name       = "PlayStation 5"
	playStation5MatchRegex = []string{`(?i)PlayStation 5`}
)

func NewPlayStation5(userAgent string) *PlayStation5 {
	return &PlayStation5{
		base{
			userAgent: userAgent,
		},
	}
}

func (p *PlayStation5) Name() string {
	return playStation5Name
}

func (p *PlayStation5) Match() bool {
	return p.match(playStation5MatchRegex)
}
