package devices

type WiiU struct {
	base
}

var (
	wiiuName       = "Nintendo WiiU"
	wiiUMatchRegex = []string{`(?i)Nintendo WiiU`}
)

func NewWiiU(userAgent string) *WiiU {
	return &WiiU{
		base{
			userAgent: userAgent,
		},
	}
}

func (w *WiiU) Name() string {
	return wiiuName
}

func (w *WiiU) Match() bool {
	return w.match(wiiUMatchRegex)
}
