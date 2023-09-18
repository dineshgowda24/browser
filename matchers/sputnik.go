
package matchers

type Sputnik struct {
	base
}

var (
	sputnikName          = "Sputnik"
	sputnikVersionRegexp = []string{`SputnikBrowser/([\d.]+)`}
	sputnikMatchRegexp   = []string{`SputnikBrowser`}
)

func NewSputnik(userAgent string) *Sputnik {
	return &Sputnik{
		base: newBase(userAgent),
	}
}

func (s *Sputnik) Name() string {
	return sputnikName
}

func (s *Sputnik) Version() string {
	return s.version(sputnikVersionRegexp, 1)
}

func (s *Sputnik) Match() bool {
	return s.match(sputnikMatchRegexp)
}
