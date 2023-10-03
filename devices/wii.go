package devices

type Wii struct {
	base
}

var (
	wiiName       = "Nintendo Wii"
	wiiMatchRegex = []string{`(?i)Nintendo Wii`}
)

func NewWii(userAgent string) *Wii {
	return &Wii{
		base{
			userAgent: userAgent,
		},
	}
}

func (w *Wii) Name() string {
	return wiiName
}

func (w *Wii) Match() bool {
	return w.match(wiiMatchRegex)
}
