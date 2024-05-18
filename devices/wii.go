package devices

type Wii struct {
	p Parser
}

var (
	wiiName       = "Nintendo Wii"
	wiiMatchRegex = []string{`(?i)Nintendo Wii`}
)

func NewWii(p Parser) *Wii {
	return &Wii{
		p: p,
	}
}

func (w *Wii) Name() string {
	return wiiName
}

func (w *Wii) Match() bool {
	return w.p.Match(wiiMatchRegex)
}
