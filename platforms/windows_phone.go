package platforms

type WindowsPhone struct {
	p Parser
}

var (
	windowsPhoneName          = "Windows Phone"
	windowsPhoneVersionRegexp = []string{`Windows Phone ([\d.]+)`}
	windowsPhoneMatchRegexp   = []string{`Windows Phone`}
)

func NewWindowsPhone(p Parser) *WindowsPhone {
	return &WindowsPhone{
		p: p,
	}
}

func (w *WindowsPhone) Name() string {
	return windowsPhoneName
}

func (w *WindowsPhone) Version() string {
	return w.p.Version(windowsPhoneVersionRegexp, 1, "")
}

func (w *WindowsPhone) Match() bool {
	return w.p.Match(windowsPhoneMatchRegexp)
}
