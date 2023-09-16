package matchers

type Konqueror struct {
	base
}

var (
	konquerorName          = "Konqueror"
	konquerorVersionRegexp = []string{`(?i)Konqueror/([\d.]+)`}
	konquerorMatchRegex    = []string{`(?i)Konqueror`}
)

func NewKonqueror(userAgent string) *Konqueror {
	return &Konqueror{
		base: newBase(userAgent),
	}
}

func (k *Konqueror) Name() string {
	return konquerorName
}

func (k *Konqueror) Version() string {
	return k.version(konquerorVersionRegexp, 1)
}

func (k *Konqueror) Match() bool {
	return k.match(konquerorMatchRegex)
}
