package platforms

type Windows struct {
	p Parser
}

var (
	windowsName          = "Windows"
	windowsVersionRegexp = []string{`Windows NT\s*([0-9_.]+)?`}
	windowsMatchRegexp   = []string{`Windows`}
)

func NewWindows(p Parser) *Windows {
	return &Windows{
		p: p,
	}
}

func (w *Windows) Name() string {
	return windowsName
}

func (w *Windows) Version() string {
	return w.p.Version(windowsVersionRegexp, 1, "0")
}

func (w *Windows) Match() bool {
	return w.p.Match(windowsMatchRegexp)
}
