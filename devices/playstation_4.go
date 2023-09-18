package devices

type PlayStation4 struct {
	base
}

var (
	playStation4Name       = "PlayStation 4"
	playStation4MatchRegex = []string{`(?i)PLAYSTATION 4`}
)

func NewPlayStation4(userAgent string) *PlayStation4 {
	return &PlayStation4{
		base{
			userAgent: userAgent,
		},
	}
}

func (p *PlayStation4) Name() string {
	return playStation4Name
}

func (p *PlayStation4) Match() bool {
	return p.match(playStation4MatchRegex)
}
