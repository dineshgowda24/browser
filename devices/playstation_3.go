package devices

type PlayStation3 struct {
	base
}

var (
	playStation3Name       = "PlayStation 3"
	playStation3MatchRegex = []string{`(?i)PLAYSTATION 3`}
)

func NewPlayStation3(userAgent string) *PlayStation3 {
	return &PlayStation3{
		base{
			userAgent: userAgent,
		},
	}
}

func (p *PlayStation3) Name() string {
	return playStation3Name
}

func (p *PlayStation3) Match() bool {
	return p.match(playStation3MatchRegex)
}
