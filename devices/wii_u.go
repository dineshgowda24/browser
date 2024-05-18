package devices

type WiiU struct {
	p Parser
}

var (
	wiiuName       = "Nintendo WiiU"
	wiiUMatchRegex = []string{`(?i)Nintendo WiiU`}
)

func NewWiiU(p Parser) *WiiU {
	return &WiiU{
		p: p,
	}
}

func (w *WiiU) Name() string {
	return wiiuName
}

func (w *WiiU) Match() bool {
	return w.p.Match(wiiUMatchRegex)
}
