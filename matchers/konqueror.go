package matchers

type Konqueror struct {
	p Parser
}

var (
	konquerorName          = "Konqueror"
	konquerorVersionRegexp = []string{`(?i)Konqueror/([\d.]+)`}
	konquerorMatchRegex    = []string{`(?i)Konqueror`}
)

func NewKonqueror(p Parser) *Konqueror {
	return &Konqueror{
		p: p,
	}
}

func (k *Konqueror) Name() string {
	return konquerorName
}

func (k *Konqueror) Version() string {
	return k.p.Version(konquerorVersionRegexp, 1)
}

func (k *Konqueror) Match() bool {
	return k.p.Match(konquerorMatchRegex)
}
