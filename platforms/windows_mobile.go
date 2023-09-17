package platforms

type WindowsMobile struct {
	base
}

var (
	windowsMobileName        = "Windows Mobile"
	windowsMobileMatchRegexp = []string{`Windows CE`}
)

func NewWindowsMobile(userAgent string) *WindowsMobile {
	return &WindowsMobile{
		base: base{
			userAgent: userAgent,
		},
	}
}

func (w *WindowsMobile) Name() string {
	return windowsMobileName
}

func (w *WindowsMobile) Version() string {
	return "0"
}

func (w *WindowsMobile) Match() bool {
	return w.match(windowsMobileMatchRegexp)
}
