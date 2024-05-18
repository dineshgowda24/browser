package devices

type PlayStation3 struct {
	p Parser
}

var (
	playStation3Name       = "PlayStation 3"
	playStation3MatchRegex = []string{`(?i)PLAYSTATION 3`}
)

func NewPlayStation3(p Parser) *PlayStation3 {
	return &PlayStation3{
		p: p,
	}
}

func (p *PlayStation3) Name() string {
	return playStation3Name
}

func (p *PlayStation3) Match() bool {
	return p.p.Match(playStation3MatchRegex)
}
