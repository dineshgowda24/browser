package matchers

type Electron struct {
	base
}

var (
	electronName          = "Electron"
	electronVersionRegexp = []string{`Electron/([\d.]+)`}
	electronMatchRegex    = []string{`Electron`}
)

func NewElectron(userAgent string) *Electron {
	return &Electron{
		base: newBase(userAgent),
	}
}

func (e *Electron) Name() string {
	return electronName
}

func (e *Electron) Version() string {
	return e.version(electronVersionRegexp, 1)
}

func (e *Electron) Match() bool {
	return e.match(electronMatchRegex)
}
