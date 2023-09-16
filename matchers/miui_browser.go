
package matchers

type MiuiBrowser struct {
	base
}

var (
	miuiBrowserName          = "Miui Browser"
	miuiBrowserVersionRegexp = []string{`MiuiBrowser/([\d.]+)`}
	miuiBrowserMatchRegexp   = []string{`MiuiBrowser`}
)

func NewMiuiBrowser(userAgent string) *MiuiBrowser {
	return &MiuiBrowser{
		base: newBase(userAgent),
	}
}

func (m *MiuiBrowser) Name() string {
	return miuiBrowserName
}

func (m *MiuiBrowser) Version() string {
	return m.version(miuiBrowserVersionRegexp, 1)
}

func (m *MiuiBrowser) Match() bool {
	return m.match(miuiBrowserMatchRegexp)
}
