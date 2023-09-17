package platforms

type WindowsPhone struct {
	base
}

var (
	windowsPhoneName          = "Windows Phone"
	windowsPhoneVersionRegexp = []string{`Windows Phone ([\d.]+)`}
	windowsPhoneMatchRegexp   = []string{`Windows Phone`}
)

func NewWindowsPhone(userAgent string) *WindowsPhone {
	return &WindowsPhone{
		base: base{
			userAgent: userAgent,
		},
	}
}

func (w *WindowsPhone) Name() string {
	return windowsPhoneName
}

func (w *WindowsPhone) Version() string {
	return w.version(windowsPhoneVersionRegexp, 1, "")
}

func (w *WindowsPhone) Match() bool {
	return w.match(windowsPhoneMatchRegexp)
}
