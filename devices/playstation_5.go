package devices

type PlayStation5 struct {
	p Parser
}

var (
	playStation5Name       = "PlayStation 5"
	playStation5MatchRegex = []string{`(?i)PlayStation 5`}
)

func NewPlayStation5(p Parser) *PlayStation5 {
	return &PlayStation5{
		p: p,
	}
}

func (p *PlayStation5) Name() string {
	return playStation5Name
}

func (p *PlayStation5) Match() bool {
	return p.p.Match(playStation5MatchRegex)
}
