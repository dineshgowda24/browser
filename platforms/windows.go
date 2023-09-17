package platforms

type Windows struct {
	base
}

var (
	windowsName          = "Windows"
	windowsVersionRegexp = []string{`Windows NT\s*([0-9_.]+)?`}
	windowsMatchRegexp   = []string{`Windows`}
)

func NewWindows(userAgent string) *Windows {
	return &Windows{
		base: base{
			userAgent: userAgent,
		},
	}
}

func (w *Windows) Name() string {
	return windowsName
}

func (w *Windows) Version() string {
	return w.version(windowsVersionRegexp, 1, "0")
}

func (w *Windows) Match() bool {
	return w.match(windowsMatchRegexp)
}
