package matchers

type Electron struct {
	p Parser
}

var (
	electronName          = "Electron"
	electronVersionRegexp = []string{`Electron/([\d.]+)`}
	electronMatchRegex    = []string{`Electron`}
)

func NewElectron(p Parser) *Electron {
	return &Electron{
		p: p,
	}
}

func (e *Electron) Name() string {
	return electronName
}

func (e *Electron) Version() string {
	return e.p.Version(electronVersionRegexp, 1)
}

func (e *Electron) Match() bool {
	return e.p.Match(electronMatchRegex)
}
